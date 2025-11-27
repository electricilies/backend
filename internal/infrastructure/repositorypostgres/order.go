package repositorypostgres

import (
	"context"

	"backend/internal/domain"
	"backend/internal/infrastructure/repositorypostgres/sqlc"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Order struct {
	queries *sqlc.Queries
	conn    *pgxpool.Pool
}

var _ domain.OrderRepository = (*Order)(nil)

func ProvideOrder(q *sqlc.Queries, conn *pgxpool.Pool) *Order {
	return &Order{
		queries: q,
		conn:    conn,
	}
}

func (r *Order) List(ctx context.Context, params domain.OrderRepositoryListParam) (*[]domain.Order, error) {
	orderEntities, err := r.queries.ListOrders(ctx, sqlc.ListOrdersParams{
		IDs:    params.IDs,
		Offset: int32(params.Offset),
		Limit:  int32(params.Limit),
	})
	if err != nil {
		return nil, toDomainError(err)
	}

	orderIDs := make([]uuid.UUID, 0, len(orderEntities))
	statusIDs := make([]uuid.UUID, 0, len(orderEntities))
	providerIDs := make([]uuid.UUID, 0, len(orderEntities))
	for _, o := range orderEntities {
		orderIDs = append(orderIDs, o.ID)
		statusIDs = append(statusIDs, o.StatusID)
		providerIDs = append(providerIDs, o.ProviderID)
	}

	orderItems, err := r.queries.ListOrderItems(ctx, sqlc.ListOrderItemsParams{
		OrderIDs: orderIDs,
	})
	if err != nil {
		return nil, toDomainError(err)
	}

	productVariantIDs := make([]uuid.UUID, 0, len(orderItems))
	for _, item := range orderItems {
		productVariantIDs = append(productVariantIDs, item.ProductVariantID)
	}
	productVariantEntities, err := r.queries.ListProductVariants(ctx, sqlc.ListProductVariantsParams{
		IDs: productVariantIDs,
	})
	if err != nil {
		return nil, toDomainError(err)
	}
	productVariantIDProductIDMap := make(map[uuid.UUID]uuid.UUID, len(productVariantEntities))
	for _, pv := range productVariantEntities {
		productVariantIDProductIDMap[pv.ID] = pv.ProductID
	}

	orderItemsMap := make(map[uuid.UUID][]domain.OrderItem)
	for _, item := range orderItems {
		orderItemsMap[item.OrderID] = append(orderItemsMap[item.OrderID], domain.OrderItem{
			ID:               item.ID,
			ProductID:        productVariantIDProductIDMap[item.ProductVariantID],
			ProductVariantID: item.ProductVariantID,
			Quantity:         int(item.Quantity),
			Price:            numericToInt64(item.Price),
		})
	}

	statusMap, err := r.getStatusMap(ctx, statusIDs)
	if err != nil {
		return nil, err
	}

	providerMap, err := r.getProviderMap(ctx, providerIDs)
	if err != nil {
		return nil, err
	}

	orders := make([]domain.Order, 0, len(orderEntities))
	for _, o := range orderEntities {
		orders = append(orders, domain.Order{
			ID:          o.ID,
			Address:     o.Address,
			Provider:    providerMap[o.ProviderID],
			Status:      statusMap[o.StatusID],
			IsPaid:      o.IsPaid,
			CreatedAt:   o.CreatedAt.Time,
			UpdatedAt:   o.UpdatedAt.Time,
			Items:       orderItemsMap[o.ID],
			TotalAmount: numericToInt64(o.TotalAmount),
			UserID:      o.UserID,
		})
	}

	return &orders, nil
}

func (r *Order) Count(ctx context.Context, params domain.OrderRepositoryCountParam) (*int, error) {
	count, err := r.queries.CountOrders(ctx, sqlc.CountOrdersParams{
		IDs: params.IDs,
	})
	if err != nil {
		return nil, toDomainError(err)
	}
	result := int(count)
	return &result, nil
}

func (r *Order) Get(ctx context.Context, params domain.OrderRepositoryGetParam) (*domain.Order, error) {
	orderEntity, err := r.queries.GetOrder(ctx, sqlc.GetOrderParams{
		ID: params.ID,
	})
	if err != nil {
		return nil, toDomainError(err)
	}

	orderItems, err := r.queries.ListOrderItems(ctx, sqlc.ListOrderItemsParams{
		OrderID: params.ID,
	})
	if err != nil {
		return nil, toDomainError(err)
	}

	productVariantIDs := make([]uuid.UUID, 0, len(orderItems))
	for _, item := range orderItems {
		productVariantIDs = append(productVariantIDs, item.ProductVariantID)
	}
	productVariantEntities, err := r.queries.ListProductVariants(ctx, sqlc.ListProductVariantsParams{
		IDs: productVariantIDs,
	})
	if err != nil {
		return nil, toDomainError(err)
	}
	productVariantIDProductIDMap := make(map[uuid.UUID]uuid.UUID, len(productVariantEntities))
	for _, pv := range productVariantEntities {
		productVariantIDProductIDMap[pv.ID] = pv.ProductID
	}

	items := make([]domain.OrderItem, 0, len(orderItems))
	for _, item := range orderItems {
		items = append(items, domain.OrderItem{
			ID:               item.ID,
			ProductID:        productVariantIDProductIDMap[item.ProductVariantID],
			ProductVariantID: item.ProductVariantID,
			Quantity:         int(item.Quantity),
			Price:            numericToInt64(item.Price),
		})
	}

	status, err := r.queries.GetOrderStatus(ctx, sqlc.GetOrderStatusParams{
		ID: orderEntity.StatusID,
	})
	if err != nil {
		return nil, toDomainError(err)
	}

	provider, err := r.queries.GetOrderProvider(ctx, sqlc.GetOrderProviderParams{
		ID: orderEntity.ProviderID,
	})
	if err != nil {
		return nil, toDomainError(err)
	}

	order := &domain.Order{
		ID:          orderEntity.ID,
		Address:     orderEntity.Address,
		Provider:    domain.OrderProvider(provider.Name),
		Status:      domain.OrderStatus(status.Name),
		IsPaid:      orderEntity.IsPaid,
		CreatedAt:   orderEntity.CreatedAt.Time,
		UpdatedAt:   orderEntity.UpdatedAt.Time,
		Items:       items,
		TotalAmount: numericToInt64(orderEntity.TotalAmount),
		UserID:      orderEntity.UserID,
	}

	return order, nil
}

func (r *Order) Save(ctx context.Context, params domain.OrderRepositorySaveParam) error {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return toDomainError(err)
	}
	qtx := r.queries.WithTx(tx)
	defer func() { _ = tx.Rollback(ctx) }()

	status, err := qtx.GetOrderStatus(ctx, sqlc.GetOrderStatusParams{
		Name: string(params.Order.Status),
	})
	if err != nil {
		return toDomainError(err)
	}

	provider, err := qtx.GetOrderProvider(ctx, sqlc.GetOrderProviderParams{
		Name: string(params.Order.Provider),
	})
	if err != nil {
		return toDomainError(err)
	}

	err = qtx.UpsertOrder(ctx, sqlc.UpsertOrderParams{
		ID:          params.Order.ID,
		UserID:      params.Order.UserID,
		Address:     params.Order.Address,
		TotalAmount: int64ToNumeric(params.Order.TotalAmount),
		IsPaid:      params.Order.IsPaid,
		ProviderID:  provider.ID,
		StatusID:    status.ID,
		CreatedAt: pgtype.Timestamptz{
			Time:  params.Order.CreatedAt,
			Valid: true,
		},
		UpdatedAt: pgtype.Timestamptz{
			Time:  params.Order.UpdatedAt,
			Valid: true,
		},
	})
	if err != nil {
		return toDomainError(err)
	}

	err = qtx.CreateTempTableOrderItems(ctx)
	if err != nil {
		return toDomainError(err)
	}

	itemParams := make([]sqlc.InsertTempTableOrderItemsParams, len(params.Order.Items))
	for i, item := range params.Order.Items {
		itemParams[i] = sqlc.InsertTempTableOrderItemsParams{
			ID:               item.ID,
			OrderID:          params.Order.ID,
			ProductVariantID: item.ProductVariantID,
			Quantity:         int32(item.Quantity),
			Price:            int64ToNumeric(item.Price),
		}
	}
	_, err = qtx.InsertTempTableOrderItems(ctx, itemParams)
	if err != nil {
		return toDomainError(err)
	}

	err = qtx.MergeOrderItemsFromTemp(ctx)
	if err != nil {
		return toDomainError(err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return toDomainError(err)
	}

	return nil
}

func (r *Order) getStatusMap(ctx context.Context, statusIDs []uuid.UUID) (map[uuid.UUID]domain.OrderStatus, error) {
	statuses, err := r.queries.ListOrderStatuses(ctx, sqlc.ListOrderStatusesParams{
		IDs: statusIDs,
	})
	if err != nil {
		return nil, toDomainError(err)
	}
	statusMap := make(map[uuid.UUID]domain.OrderStatus, len(statuses))
	for _, s := range statuses {
		statusMap[s.ID] = domain.OrderStatus(s.Name)
	}
	return statusMap, nil
}

func (r *Order) getProviderMap(ctx context.Context, providerIDs []uuid.UUID) (map[uuid.UUID]domain.OrderProvider, error) {
	providerMap := make(map[uuid.UUID]domain.OrderProvider, len(providerIDs))
	for _, id := range providerIDs {
		provider, err := r.queries.GetOrderProvider(ctx, sqlc.GetOrderProviderParams{
			ID: id,
		})
		if err != nil {
			return nil, toDomainError(err)
		}
		providerMap[id] = domain.OrderProvider(provider.Name)
	}
	return providerMap, nil
}

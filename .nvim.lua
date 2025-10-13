vim.lsp.config("gopls", {
  settings = {
    gopls = {
      buildFlags = {
        "-tags=",
      },
    },
  },
})

vim.filetype.add({
  pattern = {
    [".*%.sql"] = "pgsql",
  },
})

vim.env.DBUI_URL = string.format(
  "postgres://%s:%s@%s:%s/%s?sslmode=disable",
  vim.env.DB_USERNAME,
  vim.env.DB_PASSWORD,
  vim.env.DB_HOST,
  vim.env.DB_PORT,
  vim.env.DB_DATABASE
)
vim.env.DBUI_NAME = "electricilies"

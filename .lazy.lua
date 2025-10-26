local basename = vim.fs.basename
local env = vim.env
local nc_available = vim.fn.executable("nc")
local psql_available = vim.fn.executable("psql")

local db_username = env.DB_USERNAME
local db_password = env.DB_PASSWORD
local db_host = env.DB_HOST
local db_port = env.DB_PORT
local db_database = env.DB_DATABASE

local psql_cmd = string.format(
  'PGPASSWORD="%s" psql -h %s -p %s -U %s -d %s -c "\\q" >/dev/null 2>&1',
  db_password,
  db_host,
  db_port,
  db_username,
  db_database
)

---@return boolean
function _G._IsDbUp()
  if nc_available == 1 then
    local nc_cmd = string.format("nc -z %s %s", db_host, db_port)
    local nc_status = os.execute(nc_cmd)
    if nc_status == 0 then
      return true
    end
  end

  if psql_available == 1 then
    local psql_status = os.execute(psql_cmd)
    if psql_status == 0 then
      return true
    end
  end

  return false
end

---@module 'lazy'
---@type LazySpec
return {
  {
    "stevearc/conform.nvim",
    ---@module 'conform'
    ---@type conform.setupOpts
    opts = {
      formatters = {
        swag_fmt = {
          command = "swag",
          args = {
            "fmt",
            "-d",
            "$FILENAME",
          },
          condition = function(_, ctx)
            return ctx.filename:match("internal/interface/api/handler") ~= nil
          end,
          stdin = false,
        },
        swag_gen = {
          command = "swag",
          args = {
            "init",
            "-g",
            "./cmd/main.go",
            "-ot",
            "go",
          },
          condition = function(_, ctx)
            return ctx.filename:match("internal/interface/api/handler") ~= nil
          end,
          stdin = false,
        },
        wire = {
          command = "wire",
          args = {
            "gen",
            "$DIRNAME",
          },
          condition = function(_, ctx)
            return basename(ctx.filename):match(".*%.wire.go") ~= nil
          end,
          stdin = false,
        },
        sqlc = {
          command = "sqlc",
          args = {
            "generate",
          },
          condition = function(_, ctx)
            return ctx.filename:match("database/.*.sql") ~= nil
          end,
        },
        atlas = {
          command = "atlas",
          args = {
            "schema",
            "apply",
            "--env=local",
            "--auto-approve",
          },
          stdin = false,
          condition = function(_, ctx)
            return _IsDbUp() and basename(ctx.filename):match("schema.sql") ~= nil
          end,
        },
      },
      formatters_by_ft = {
        go = {
          "swag_fmt",
          "swag_gen",
          "wire",
        },
        pgsql = {
          "atlas",
          "sqlc",
        },
      },
    },
    opts_extend = {
      "formatters_by_ft.go",
      "formatters_by_ft.pgsql",
    },
    optional = true,
  },
  -- {
  --   "mfussenegger/nvim-lint",
  --   opts = function()
  --     local lint = require("lint")
  --     lint.linters.sqlc = function()
  --       local bufname = vim.api.nvim_buf_get_name(0)
  --       if bufname:match("/database/.*%.sql$") and is_db_up() then
  --         ---@type lint.Linter
  --         return {
  --           name = "sqlc",
  --           cmd = "sqlc",
  --           args = { "vet" },
  --           stream = "stderr",
  --           parser = require("lint.parser").from_pattern(
  --             "^(.+): (.+: .+): (.+)$",
  --             { "file", "code", "message" },
  --             nil,
  --             {
  --               source = "sqlc",
  --               severity = vim.diagnostic.severity.WARN,
  --             }
  --           ),
  --         }
  --       end
  --       return {}
  --     end
  --     lint.linters_by_ft.pgsql = lint.linters_by_ft.pgsql or {}
  --     table.insert(lint.linters_by_ft.pgsql, "sqlc")
  --   end,
  --   optional = true,
  -- },
  {
    "kristijanhusak/vim-dadbod-ui",
    opts = function()
      env.DBUI_URL = string.format(
        "postgres://%s:%s@%s:%s/%s?sslmode=disable",
        db_username,
        db_password,
        db_host,
        db_port,
        db_database
      )
      env.DBUI_NAME = "electricilies-local"
    end,
    optional = true,
  },
  {
    "folke/which-key.nvim",
    ---@module 'which-key'
    ---@type wk.Opts
    opts = {
      spec = {
        { "<localleader>l", group = "LSP", icon = { icon = "", color = "yellow" } },
      },
    },
    opts_extend = {
      "spec",
    },
    optional = true,
  },
}

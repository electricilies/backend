local basename = vim.fs.basename
local env = vim.env

---@module 'lazy'
---@type LazySpec
return {
  {
    "stevearc/conform.nvim",
    ---@module 'conform'
    ---@type conform.setupOpts
    opts = {
      formatters = {
        swag = {
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
            return basename(ctx.filename):match("schema.sql") ~= nil
          end,
        },
      },
      formatters_by_ft = {
        go = {
          "swag",
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
  {
    "mfussenegger/nvim-lint",
    opts = function()
      local lint = require("lint")

      lint.linters.sqlc = function()
        local bufname = vim.api.nvim_buf_get_name(0)
        if bufname:match("/database/.*%.sql$") then
          ---@type lint.Linter
          return {
            name = "sqlc",
            cmd = "sqlc",
            args = { "vet" },
            stream = "stderr",
            parser = require("lint.parser").from_pattern(
              "^(.+): (.+: .+): (.+)$",
              { "file", "code", "message" },
              nil,
              {
                source = "sqlc",
                severity = vim.diagnostic.severity.WARN,
              }
            ),
          }
        end
        return {}
      end

      lint.linters_by_ft.pgsql = lint.linters_by_ft.pgsql or {}
      table.insert(lint.linters_by_ft.pgsql, "sqlc")
    end,
    optional = true,
  },
  {
    "kristijanhusak/vim-dadbod-ui",
    opts = function()
      env.DBUI_URL = string.format(
        "postgres://%s:%s@%s:%s/%s?sslmode=disable",
        env.DB_USERNAME,
        env.DB_PASSWORD,
        env.DB_HOST,
        env.DB_PORT,
        env.DB_DATABASE
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

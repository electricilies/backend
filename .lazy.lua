local basename = vim.fs.basename

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
        },
      },
      formatters_by_ft = {
        go = {
          "swag",
          "wire",
        },
        pgsql = {
          "sqlc",
        },
      },
    },
    opts_extend = {
      "formatters_by_ft.go",
      "formatters_by_ft.pgsql",
    },
  },
  {
    "mfussenegger/nvim-lint",
    opts = function()
      local lint = require("lint")

      lint.linters.sqlc = {
        name = "sqlc",
        cmd = "sqlc",
        args = { "vet" },
        stream = "stderr",
        ignore_exitcode = true,
        parser = require("lint.parser").from_pattern("^(.+): (.+: .+): (.+)$", { "file", "code", "message" }, nil, {
          source = "sqlc",
          severity = vim.diagnostic.severity.WARN,
        }),
      }

      lint.linters_by_ft.pgsql = lint.linters_by_ft.pgsql or {}
      table.insert(lint.linters_by_ft.pgsql, "sqlc")
    end,
  },
}

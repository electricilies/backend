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
          command = "go",
          args = {
            "tool",
            "swag",
            "fmt",
            "-d",
            "$FILENAME",
          },
          stdin = false,
        },
        wire = {
          command = "go",
          args = {
            "tool",
            "wire",
            "gen",
            "$RELATIVE_FILEPATH",
          },
          condition = function(_, ctx)
            return basename(ctx.filename) == "wire.go"
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
        postgresql = {
          "sqlc",
        },
        sql = {
          "sqlc",
        },
      },
    },
    opts_extend = {
      "formatters_by_ft.go",
      "formatters_by_ft.postgresql",
      "formatters_by_ft.sql",
    },
  },
}

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
        sqlc_gen = {
          command = "sqlc",
          args = {
            "generate",
          },
        },
      },
      formatters_by_ft = {
        go = {
          "swag",
        },
        postgresql = {
          "sqlc_gen",
        },
        sql = {
          "sqlc_gen",
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

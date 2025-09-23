---@module 'lazy'
---@type LazySpec
return {
  {
    "stevearc/conform.nvim",
    ---@module 'conform'
    ---@type conform.setupOpts
    opts = {
      formatters = {
        sqlc_gen = {
          command = "sqlc",
          args = {
            "generate",
          },
        },
      },
      formatters_by_ft = {
        sql = {
          "sqlc_gen",
        },
        postgresql = {
          "sqlc_gen",
        },
      },
    },
    opts_extend = {
      "formatters_by_ft.sql",
      "formatters_by_ft.postgresql",
    },
  },
}

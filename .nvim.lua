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
  extension = {
    tf = "terraform",
  },
})

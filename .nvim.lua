vim.lsp.config("gopls", {
  settings = {
    gopls = {
      buildFlags = {
        "-tags=",
      },
    },
  },
})

-- if vim.lsp.config["swaggo_ls"] then
--  vim.lsp.enable("swaggo_ls")
-- end

vim.filetype.add({
  pattern = {
    [".*%.sql"] = "pgsql",
  },
  extension = {
    tf = "terraform",
  },
})

local map = vim.keymap.set

vim.lsp.config("gopls", {
  settings = {
    gopls = {
      buildFlags = {
        "-tags=integration",
      },
    },
  },
})

map("n", "<localleader>lb", function()
  vim.ui.select({
    "none",
    "integration",
    "wireinject",
  }, {
    prompt = "Select gopls build flags:",
  }, function(flag)
    vim.lsp.config["gopls"] = {
      settings = {
        buildFlags = {
          "-tags" .. flag,
        },
      },
    }
    for _, client in ipairs(vim.lsp.client()) do
      if client.name == "gopls" then
        vim.lsp.stop_client(client, true)
      end
    end
    vim.lsp.start(vim.lsp.config["gopls"])
  end)
end, { desc = "LSP | Switch buildFlags", silent = true })

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

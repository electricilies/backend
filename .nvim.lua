local map = vim.keymap.set
local lsp = vim.lsp

lsp.config("gopls", {
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
    if not flag then
      return
    end
    lsp.config["gopls"] = {
      settings = {
        buildFlags = {
          "-tags=" .. flag ~= "none" and flag or "",
        },
      },
    }
    local clients = lsp.get_clients({ name = "gopls" })
    lsp.stop_client(clients, true)
    lsp.start(lsp.config["gopls"])
  end)
end, { desc = "LSP | Switch buildFlags", silent = true })

-- if lsp.config["swaggo_ls"] then
--  lsp.enable("swaggo_ls")
-- end

vim.filetype.add({
  pattern = {
    [".*%.sql"] = "pgsql",
  },
  extension = {
    tf = "terraform",
  },
})

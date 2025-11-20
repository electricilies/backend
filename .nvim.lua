local map = vim.keymap.set
local lsp = vim.lsp
local g = vim.g

map("n", "<localleader>g", function()
  g.dev_no_gen = not g.dev_no_gen
  vim.notify("Dev Gen is " .. (g.dev_no_gen and "Disabled" or "Enabled"))
end, { desc = "General | Toggle Dev Gen", silent = true })

lsp.config("gopls", {
  settings = {
    gopls = {
      buildFlags = {
        "-tags",
        "integration",
      },
    },
  },
})

lsp.config("postgres_lsp", {
  root_dir = function(_, on_dir)
    if _IsDbUp() then
      on_dir()
    end
  end,
})

map("n", "<localleader>lb", function()
  vim.ui.select({
    "none",
    "integration",
    "wireinject",
    "integration,wireinject",
  }, {
    prompt = "Select gopls build tag",
  }, function(tag)
    if not tag then
      return
    end
    local clients = lsp.get_clients({ name = "gopls" })
    lsp.stop_client(clients, true)
    lsp.config.gopls = {
      settings = {
        gopls = {
          buildFlags = tag ~= "none" and {
            "-tags",
            tag,
          } or {},
        },
      },
    }
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

vim.api.nvim_create_autocmd("BufWritePost", {
  pattern = "*.go",
  callback = function(args)
    local filename = args.file
    if g.dev_no_gen or filename:match("internal/repository/.*%.go") == nil then
      return
    end
    vim.fn.jobstart("mockery", {
      cwd = vim.fn.getcwd(),
      on_exit = function(_, exit_code)
        if exit_code ~= 0 then
          vim.notify("Mockery: Failed to generate mocks", vim.log.levels.ERROR)
        end
      end,
    })
  end,
  desc = "Generate mocks with mockery on save",
})

# lsp-minusOne

A Language Server built for the education purpose of understanding **WHAT** 
lsp is and **HOW** it works.

It would not do anything special for any particular language, it is geared towards
helping you understand what your tools **do**.

## Setting up Neovim with your LSP

1. Access the LSP client
2. Handle errors if any
3. Attach the LSP to a buffer (here, we are doing markdowns)

Filename : `load_test_lsp.lua`
```lua
local client = vim.lsp.start_client {
  name = "lsp-minusOne",
  cmd = {
    "absolute-path-to-lsp-binary",
  },
  on_attach = require("rk.lsp").on_attach,
}

if not client then
  vim.notify "Client not initialized!"
  return
end

vim.api.nvim_create_autocmd("FileType", {
  pattern = "markdown"
  callback = function()
    vim.lsp.buf_attach_client(0, client)   
  end,
})
```

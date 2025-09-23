atlas-hash env="local":
  atlas migrate hash --env {{env}}

atlas-gen-migration env="local":
  atlas migrate diff --env {{env}}

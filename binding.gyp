{
  "targets": [
    {
      "target_name": "express",
      "sources": [ "src/express.cc" ],
      "include_dirs": ["src"],
      "libraries": [
        "<(module_root_dir)/libexpress.a"
      ]
    }
  ]
}
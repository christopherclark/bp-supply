#!/usr/bin/env ruby
require "webrick"

server = WEBrick::HTTPServer.new(:Port => (ENV['PORT'] || 8080))
server.mount_proc '/' do |req, res|
  if req.path == '/supply_config'
    res.body = File.read("#{ENV['DEPS_DIR']}/0/config.yml")
  elsif req.path === '/'
    res.body = `jane`
  else
    res.body = 'Not Found'
  end
end

trap("INT") { server.shutdown }
server.start

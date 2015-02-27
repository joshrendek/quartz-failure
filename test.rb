require 'bundler/setup'
require 'quartz'
client = Quartz::Client.new(file_path: 'example.go')
res = client[:foos].call('Test', 'Num' => 1)
puts res

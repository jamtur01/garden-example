require 'bundler'
require 'sinatra'

port = ENV['PORT'] || 8080
puts "STARTING SINATRA on port #{port}"
set :port, port
set :bind, '0.0.0.0'

get '/mate' do
  "Mate!"
end

get '/yo' do
  redirect "http://yo-service/yo"
end


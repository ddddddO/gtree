require 'sinatra'
require 'sinatra/reloader'

# refs
# sinatra: http://sinatrarb.com/intro-ja.html
# bundle init/Gemfile/bundle exec ~: https://qiita.com/oshou/items/6283c2315dc7dd244aef
# sinatra オートリロード: https://qiita.com/izumin5210/items/cd2f9f48fbe1fdcaf628

#enable :sessions
use Rack::Session::Pool,
    :key => 'test.ss',  # default session key name: rack.session
    :expire_after => 20 # 20s

#set :session_secret, 'xxxx'

puts 'launch app server'

get '/' do
    "session value = " << session[:v].inspect
end

get '/:value' do
    v = params['value']
    session[:v] = v
    v
end

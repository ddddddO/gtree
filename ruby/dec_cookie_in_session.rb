require 'pp'
require 'base64'
require 'uri'
require 'openssl'

# ref: https://riocampos-tech.hatenablog.com/entry/20140616/private_study_about_rack_session_or_cookie
# の「cookieの中身を確認」から下

# COOKIES rack.session(HTTP_COOKIE) の値(サーバ側)
session_in_cookie = "BAh7CkkiD3Nlc3Npb25faWQGOgZFVEkiRTZkNzcxMTU0NDBmNWJjNWMyODFj%0AMTNkOThiZTJhZTc0ZjFlZmE0MTMzMDk2ZmM3YmE1MTMyNjUyMmU5ZWUxYWQG%0AOwBGSSIJY3NyZgY7AEZJIjE2Q0ZiS1hiWnhIcnJlM1NzQm1sdGFxV2RIQnR5%0ANjZrb2NMZ0xLRUxVamhZPQY7AEZJIg10cmFja2luZwY7AEZ7BkkiFEhUVFBf%0AVVNFUl9BR0VOVAY7AFRJIi1kZTEzMjYyNGE2YmIyODliNzU0M2E5Mzk2NmUy%0ANTY2MWE3MmUzZmJiBjsARkkiDHVzZXJfaWQGOwBGSSIGMQY7AFRJIg5tZXNz%0AYWdlZWUGOwBGSSITSGVsbG8gc2Vzc2lvbiEGOwBU%0A--e507997ee379d1d63200aad5645cad96b6d492f1"
session_base64, digest = URI.decode(session_in_cookie).split("--")

pp session_base64
puts ''

pp Marshal.load(Base64.decode64(session_base64))
#{"session_id"=>
#    "6d77115440f5bc5c281c13d98be2ae74f1efa4133096fc7ba51326522e9ee1ad",
#   "csrf"=>"6CFbKXbZxHrre3SsBmltaqWdHBty66kocLgLKELUjhY=",
#   "tracking"=>{"HTTP_USER_AGENT"=>"de132624a6bb289b7543a93966e25661a72e3fbb"},
#   "user_id"=>"1",
#   "messageee"=>"Hello session!"}

# 以下ブラウザ(Chrome)で保持しているCookie　と合致！
# BAh7CkkiD3Nlc3Npb25faWQGOgZFVEkiRTZkNzcxMTU0NDBmNWJjNWMyODFj%0AMTNkOThiZTJhZTc0ZjFlZmE0MTMzMDk2ZmM3YmE1MTMyNjUyMmU5ZWUxYWQG%0AOwBGSSIJY3NyZgY7AEZJIjE2Q0ZiS1hiWnhIcnJlM1NzQm1sdGFxV2RIQnR5%0ANjZrb2NMZ0xLRUxVamhZPQY7AEZJIg10cmFja2luZwY7AEZ7BkkiFEhUVFBf%0AVVNFUl9BR0VOVAY7AFRJIi1kZTEzMjYyNGE2YmIyODliNzU0M2E5Mzk2NmUy%0ANTY2MWE3MmUzZmJiBjsARkkiDHVzZXJfaWQGOwBGSSIGMQY7AFRJIg5tZXNz%0AYWdlZWUGOwBGSSITSGVsbG8gc2Vzc2lvbiEGOwBU%0A--e507997ee379d1d63200aad5645cad96b6d492f1
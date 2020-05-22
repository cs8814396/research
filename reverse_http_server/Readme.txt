install wrk:

sudo apt-get install build-essential libssl-dev git -y
git clone https://github.com/wg/wrk.git 
cd wrk
sudo make
# move the executable to somewhere in your PATH, ex:
sudo cp wrk /usr/local/bin



-------------------------






test result:


net/http reverse_server request net/http server and reverse locally

 wrk -t1 -c4 -d20s 'http://127.0.0.1:14001'
Running 20s test @ http://127.0.0.1:14001
  1 threads and 4 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   556.46us  489.62us  12.77ms   90.99%
    Req/Sec     8.10k   776.99     9.38k    74.63%
  161926 requests in 20.10s, 24.40MB read
Requests/sec:   8056.11
Transfer/sec:      1.21MB



net/http reverse_server request net/http server and reverse locally
sudo tcpdump -Annls0 -iany port 8000
curl "http://127.0.0.1:14001"  请求内容无法抓取，是https

curl -v --http2 http://localhost:8000 校验
require 'unirest'
require 'pp'

def getID
	File.read("id.txt")
end

def setID(id)
	File.write("id.txt", id)
end

def make_query(params)
	params.map { |e| e.join("=") }.join("&")
end

def send(method, params)
	Unirest.get("https://api.vk.com/method/#{method}?#{make_query(params)}&v=5.50&access_token=#{@token}").body["response"]
end

def handle_request(msg)
	pp [msg["id"], msg["user_id"], msg["body"]]
	message = msg["body"].split(" ")
	if message[0][0] == "!"
		p "Bot command spotted"
		# send("messages.send", [["chat_id", @chat_id], ["message", "test"]])
	end
end

@token = "3ef577c0fad1cbb7d7d82614d53d3b81f78e5425edfd4abac03ea757d4c5f96a1b4a43ff25e3d0b5eb7b5"
@chat_id = 19
Unirest.timeout(5)

setID(send("messages.get", [["count", 1]])["items"].first["id"])

loop do
	begin
		# res = Unirest.get("https://api.vk.com/method/messages.get?count=100&last_message_id=#{getID}&out=0&filters=0&v=5.50&access_token=#{@token}")
		res = send("messages.get", [["count", 100], ["last_message_id", getID], ["filters", 0]])
		res = res["items"].reverse
		res.each { |e| 
			handle_request(e) if e["chat_id"] == @chat_id
		}
		setID(res.last["id"]) if !res.last.nil? && res.size > 0
		p Time.now.to_i
		sleep(5)
	rescue => e
		pp e.inspect
	end
end
class BroadcastChangesJob < ApplicationJob
	queue_as :default

	def perform
		ActionCable.server.broadcast 'DtbChannel', "new"
	end
end

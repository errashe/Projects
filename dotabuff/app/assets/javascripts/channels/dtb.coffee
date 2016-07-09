App.dtb = App.cable.subscriptions.create "DtbChannel",

	connected: ->
		# console.log "c"

	disconnected: ->
		# console.log "d"

	received: (data) ->
		console.log data

	speak: ->
		@perform 'speak'

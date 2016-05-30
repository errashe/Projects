var canvas, ctx, keyState={};

function name() {
	return document.querySelector("input[name=name]").value;
}

function draw() {
	ctx.clearRect(0,0,500,500);

	var players = Players.find().fetch();
	for(var i=0; i<players.length; i++) {
		var p = players[i];
		ctx.fillRect(p.x, p.y, 10, 10);
	}

	if (keyState[37]) {
		Meteor.call("movePlayer", name(), {x: -5});
	}
	if (keyState[38]) {
		Meteor.call("movePlayer", name(), {y: -5});
	}
	if (keyState[39]) {
		Meteor.call("movePlayer", name(), {x: 5});
	}
	if (keyState[40]) {
		Meteor.call("movePlayer", name(), {y: 5});
	}
}

function keyDown(e) {
	console.log(e);
	if (e.keyCode == 37) {
		Meteor.call("movePlayer", name(), {x: -5});
	} else if(e.keyCode == 38) {
		Meteor.call("movePlayer", name(), {y: -5});
	} else if(e.keyCode == 39) {
		Meteor.call("movePlayer", name(), {x: 5});
	} else if (e.keyCode == 40) {
		Meteor.call("movePlayer", name(), {y: 5});
	}
}

Template.game.rendered = function() {
	Meteor.subscribe("players");

	window.addEventListener('keydown',function(e){
		keyState[e.keyCode || e.which] = true;
	},true);
	window.addEventListener('keyup',function(e){
		keyState[e.keyCode || e.which] = false;
	},true);

	canvas = document.querySelector("#can");
	canvas.width = 500;
	canvas.height = 500;

	ctx = canvas.getContext("2d");
	setInterval(draw, 1000/60);
}

Template.game.events({
	'click input[name=reg]': function(e) {
		Meteor.call('regPlayer', name());
	}
});
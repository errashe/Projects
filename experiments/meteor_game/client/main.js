import { Template } from 'meteor/templating';
import { ReactiveVar } from 'meteor/reactive-var';

import './main.html';

var canvas, ctx, player;
Meteor.subscribe("players");

Template.game.events({
	'click #createNew': function(e) {
		player = ""+Date.now();
		document.querySelector("input").value = player;
		Meteor.subscribe("players", player);
		Meteor.call("addPlayer", player);
	},
	'change input': function(e) {
		player = e.target.value;
		Meteor.subscribe("players", player);
	},
	'click #c': function(e) {
		console.log("qwe");
		Meteor.call("move", player, 10);
	}
});

Template.game.helpers({
	player: function() {
		return player;
	}
});

function game() {
	var pl = Players.find({name: player}).fetch()[0];
	if (pl != undefined) {
		ctx.clearRect(0, 0, 500, 500);
		ctx.strokeRect(pl.x, pl.y, 10, 10);
	}

}

Template.game.rendered = function() {
	canvas = document.querySelector("#c");
	ctx = canvas.getContext('2d');

	setInterval(game, 1000/60);

}


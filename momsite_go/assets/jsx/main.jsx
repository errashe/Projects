const { Route, Router, Link, IndexLink, browserHistory, IndexRoute } = ReactRouter;

class App extends React.Component {
	render() {
		return (
			<div>
				<TopMenu />
				{this.props.children}
			</div>
		);
	}
}

class NavItem extends React.Component {
	render() {
		const isActive = this.context.router.isActive(this.props.to, true);
		const className = isActive ? "uk-active" : "";

		return (
			<li className={className}>
				<Link {...this.props}>{this.props.children}</Link>
			</li>
		);
	}
}

NavItem.contextTypes = {
	router: React.PropTypes.func.isRequired
}

class TopMenu extends React.Component {
	render() {
		return (
			<nav className="uk-navbar">
				<ul className="uk-navbar-nav">
					<NavItem to='/'>Home</NavItem>
					<NavItem to='/a'>A</NavItem>
				</ul>
			</nav>
		);
	}
}

class MainBox extends React.Component {
	render() {
		return (
			<div>
				<h1>Hello!</h1>
			</div>
		);
	}
}

class CommentBox extends React.Component {
	render() {
		return (
			<div className="commentBox">
				<h1>Comments</h1>
			</div>
		);
	}
}

ReactDOM.render((
	<div className="container">
		<Router history={browserHistory}>
			<Route path="/" component={App}>
				<IndexRoute component={MainBox} />
				<Route path="/test" component={CommentBox} />
			</Route>
		</Router>
	</div>
), document.getElementById("app"))
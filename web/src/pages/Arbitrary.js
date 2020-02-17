import React, { Component } from 'react';
import { Link } from 'react-router-dom'

class Arbitrary extends Component {
    constructor(props) {
        super(props);
        this.state = { apiResponse: "" };
    }

    callAPI() {
        fetch("http://localhost:8080/arbitrary")
            .then(res => res.text())
            .then(res => this.setState({ apiResponse: res }))
            .catch(err => err);
    }

    componentDidMount() {
        this.callAPI();
    }

    render() {
        return (
	    <div className="container">
              <p>{this.state.apiResponse}</p>
	      <Link to="/">Home</Link>
            </div>
        );
    }
}

export default Arbitrary;

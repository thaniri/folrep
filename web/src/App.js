import React from 'react'
import { Route, Switch } from 'react-router-dom'
import HomePage from './pages/HomePage'
import Arbitrary from './pages/Arbitrary'

export default function App() {
  return (
    <Switch>
      <Route exact path="/" component={HomePage} />
      <Route exact path="/arbitrary" component={Arbitrary} />
    </Switch>
  )
}
//import React, { Component } from 'react';
////import logo from './logo.svg';
////import './App.css';
//
//class App extends Component {
//    constructor(props) {
//        super(props);
//        this.state = { apiResponse: "" };
//    }
//
//    callAPI() {
//        fetch("http://localhost:8080/")
//            .then(res => res.text())
//            .then(res => this.setState({ apiResponse: res }))
//            .catch(err => err);
//    }
//
//    componentDidMount() {
//        this.callAPI();
//    }
//
//    render() {
//        return (
//		<p className="App-intro">{this.state.apiResponse}</p>
//        );
//    }
//}
//
//export default App;

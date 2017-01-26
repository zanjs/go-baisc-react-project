import React, { Component } from 'react';
import logo from './logo.svg';
import axios from 'axios';
import './App.css';

class App extends Component {

  state = {
    todos: []
  }
  componentWillMount() {
    var todos = axios.get('http://localhost:8000/services/todos', {});

    todos.then(content => {
      this.setState({ todos: content.data })
    }).catch(error => {
      console.log(error);
    });

  }

  render() {
    return (
      <div className="App">
        <div className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h2>React + Go</h2>
        </div>
        <div className="App-intro">
          The list of TODO items is coming from the backend implemented in Goa

            <h3>TODO</h3>
            <div></div>
            {
              this.state.todos.map((data, index) => {
                return (<div key={index}>{data.name} -- {data.completed ? 'Completed' : 'Not yet completed'}  <br /> <br /></div>)
              })
            }

        </div>
      </div>
    );
  }
}

export default App;

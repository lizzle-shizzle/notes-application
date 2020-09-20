import React from 'react';
import './App.css';
import Notes from './components/notes';
import { Col, Row } from 'reactstrap';

class App extends React.Component {
  render() {
    return (
      <div className="App">
        <br></br>
        <Row>
          <Col xs="12">
            <Notes></Notes>
          </Col>
        </Row>
      </div>
    );
  }
}

export default App;

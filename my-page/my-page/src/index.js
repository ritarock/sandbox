import React from 'react';
import ReactDOM from 'react-dom';
// import './index.css';

class Contents extends React.Component {
  render() {
    return (
      <div>
        contents
      </div>
    );
  }
}

class Header extends React.Component {
  render() {
    return (
      <div>
        my page
      </div>
    );
  }
}

class Main extends React.Component {
  render() {
    return (
      <div>
        <Header />
        <Contents />
      </div>
    );
  }
}

// ========================================

ReactDOM.render(
  <Main />,
  document.getElementById('root')
);


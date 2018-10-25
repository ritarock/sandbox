import React from 'react';
import ReactDOM from 'react-dom';
import Footer from './components/Footer';
import Header from './components/Header';
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

class Main extends React.Component {
  render() {
    return (
      <div>
        <Header />
        <Contents />
        <Footer />
      </div>
    );
  }
}

// ========================================

ReactDOM.render(
  <Main />,
  document.getElementById('root')
);


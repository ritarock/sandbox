import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import StateForm from './StateForm.tsx'
import FormTextarea from './FormTextarea.tsx'
import StateTodo from './StateTodo.tsx'
import FormBasic from './FormBasic.tsx'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <App />
    <StateForm />
    <FormTextarea />
    <StateTodo />
    <FormBasic />
  </React.StrictMode>,
)

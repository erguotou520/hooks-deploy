import React from 'react'
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'

import HomePage from './pages/Home'
import CreatePage from './pages/Product/Create'
import EditPage from './pages/Product/Edit'

const RouterComponent = () => {
  return (
    <Router>
      <Switch>
        <Route path="/edit/:id" component={EditPage} />
        <Route path="/create" component={CreatePage} />
        <Route path="/" component={HomePage} exact />
      </Switch>
    </Router>
  )
}

export default RouterComponent

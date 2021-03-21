import './App.css';
import {
  BrowserRouter as Router1,
  Switch,
  Route,
  Link
} from "react-router-dom";
import { SkillsList } from './components/Skills'
import { LanguageList } from './components/Languages'
import { Home } from './components/Home'
import { Container } from 'react-bootstrap'
import React, {createContext, useEffect, useContext, useReducer, useState} from 'react';

const initialArgs = {};
const RouterContext = createContext(initialArgs);

function RouterStore({children}) {
  const [page, setPage] = useState();
  const [params, setParams] = useState({});
  const updateParams = (newParams) => {
    setParams({...params, ...newParams})
  }
  const value = {page, setPage, params, updateParams};
  return (
    <RouterContext.Provider value={value}>
      {children}
    </RouterContext.Provider>
  )
}

function Router({children}) {
  const { page } = useContext(RouterContext);
  const pageToDisplay = children.find((c) => {
    return !!page && c.props.pagename === page
  })
  if(!!pageToDisplay){
    return pageToDisplay
  }
  return (
    <p>Página não encontrada</p>
  )
}
function Page2(){
  const {params} = useContext(RouterContext);
  return (
    <p>Params: {!!params.test && params.test}</p>
  )
}

function App() {
  return (
    <RouterStore>
      <NavBar></NavBar>
      <Router>
        <p pagename="page1">PAGE 1</p>
        <Page2 pagename="page2"></Page2>
        <p pagename="page3">PAGE 3</p>
        <p pagename="page4">PAGE 4</p>
      </Router>
    </RouterStore>
  );
}
function NavBar (){
  const {setPage, updateParams} = useContext(RouterContext);
  return (
    <>
      <button onClick={() => setPage("page1")}>Page1</button>
      <button onClick={() => setPage("page2")}>Page2</button>
      <button onClick={() => setPage("page3")}>Page3</button>
      <button onClick={() => setPage("page4")}>Page4</button>
      <button onClick={() => updateParams({test: "PARAMS OK!"})}>set params</button>
      <button onClick={() => updateParams({test: ""})}>clean params</button>
    </>
  )
}
function App1() {
  return (
    <Router>
      <div>
        <ul>
          <li>
            <Link to="/">Home</Link>
          </li>
          <li>
            <Link to="/library">Library</Link>
          </li>
          <li>
            <Link to="/library/languages">Languages</Link>
          </li>
        </ul>
        <hr />
        <Switch>
          <Container>
            <Route exact path="/">
              <Home />
            </Route>
            <Route exact path="/library">
              <SkillsList />
            </Route>
            <Route path="/library/languages">
              <LanguageList />
            </Route>
          </Container>
        </Switch>
      </div>
    </Router>
  );
}

export default App;

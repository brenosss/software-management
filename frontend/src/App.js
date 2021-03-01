import './App.css';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";
import { SkillsList } from './components/Skills'
import { LanguageList } from './components/Languages'
import { Home } from './components/Home'
import { Container } from 'react-bootstrap'


function App() {
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

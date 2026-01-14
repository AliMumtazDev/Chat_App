import "./App.css";
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Signup from "./pages/signup";
import Login from "./pages/login";
import Main from "./components/main";
function App() {
  return (
    <Router>
      <Routes>
        <Route path="/signup" element={<Signup />} />
        <Route path="/login" element={<Login />} />
        <Route path="/chat" element={<Main />} />
        <Route path="/" element={<Login />} /> 
      </Routes>
    </Router>
  );
}

export default App;

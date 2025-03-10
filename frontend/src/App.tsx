import { Route, Routes } from "react-router-dom";
import Home from "./Routes/Home/Home";
import Layout from "./Layout";
import Register from "./Routes/Register/Components/Register";
import RegisterRefer from "./Routes/Register/Components/RegisterRefer";
import AuthProvider from "./Protected/AuthProvider";
import ErrorPage from "./Routes/ErrorPage/ErrorPage";
import Login from "./Routes/Login/Login";

function App() {
  return (
    <AuthProvider>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Home />} />
          <Route
            path="register"
            element={
              <RegisterRefer submitColor="bg-indigo-800" submitText="SUBMIT" />
            }
          />
          <Route path="/register/:id" element={<Register />} />
          <Route path="/test-route" element={<Login />} /> 
          <Route path="*" element={<ErrorPage />} />
        </Route>
        
      </Routes>
    </AuthProvider>
  );
}

export default App;

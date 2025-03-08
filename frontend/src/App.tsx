import { Route, Routes } from "react-router-dom";
import Home from "./Routes/Home/Home";
import Layout from "./Layout";
import Register from "./Routes/Register/Components/Register";
import RegisterRefer from "./Routes/Register/Components/RegisterRefer";
import AuthProvider from "./Protected/AuthProvider";

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

        </Route>
      </Routes>
    </AuthProvider>
  );
}

export default App;

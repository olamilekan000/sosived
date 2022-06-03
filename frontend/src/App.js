import axios from 'axios'
import {  useState } from 'react';
import { toast } from 'react-toastify';
import Layout from './components/Layout';
import LogIn from './components/Login';

import SendButton from "./components/sendButton";

function App() {

  const [isLoggedin, setIsLoggedIn] = useState(true)
  const [isVerifying, setVerifying] = useState(false)
  const [payload, setPayload] = useState({})

  const handleChange = (e) => {
    setPayload({
      ...payload,
      [e.target.name]: e.target.value
    })
  }

  const handleToast = ({message, color}) => toast(message, {
    style: {
      backgroundColor: color,
      color:"white"
    },
    progress: false
  })

  useState(() => {

    setIsLoggedIn(false)
    setVerifying(true)

    const token = localStorage.getItem("token")
    if(!token) {
      setVerifying(false)
      return 
    }

    (async () => {
      try {
        const res = await axios.request({
          url: "http://localhost:8989/auth/verify",
          method: "GET",
          data: payload,
          headers: {
            Authorization: `Bearer ${token}`,
          }        
        })     
  
        if(res.statusText === 'OK'){
          setIsLoggedIn(true)
        }
        setVerifying(false)        
      } catch (error) {
        setVerifying(false)
        localStorage.clear()
        setIsLoggedIn(false)
      }
    })()
  }, [])

  const sendMessage = async (e) => {
    e.preventDefault()
    try {
  
      const res = await axios.request({
        url: "http://localhost:8188/send-message",
        method: "POST",
        data: payload,
        headers: {
          Authorization: `Bearer ${localStorage.getItem("token")}`,
        }
      })

      setIsLoggedIn(true)

      handleToast({message: res?.data?.message, color: "green"})
    } catch (error) {
      setIsLoggedIn(false)
      handleToast({message: error?.response?.data?.message, color: "tomato"})
    }
  }

  const login = async (e) => {
    e.preventDefault()
    try {

      if(!payload.username || !payload.password){
        handleToast({message: "Please enter all required fields.", color: "tomato"})
        return
      }
  
      const res = await axios.request({
        url: "http://localhost:8989/auth/login", //base url should come from enviroment
        method: "POST",
        data: payload,
      })

      localStorage.setItem("token", res.data.data.token)

      handleToast({message: res.data.message, color: "green"})
      setIsLoggedIn(true)

    } catch (error) {
      setIsLoggedIn(false)
      handleToast({message: error?.response?.data?.message, color:"tomato"})
    }
  }

  


  return (
    <div className="App">
      <Layout>
        {
          (isLoggedin && !isVerifying )&&(
            <SendButton sendMessage={sendMessage}/>
          )
        }

        {
          (!isLoggedin && !isVerifying) && (
            <LogIn handleChange={handleChange} login={login}/>
          )
        }
      </Layout>
    </div>
  );
}

export default App;

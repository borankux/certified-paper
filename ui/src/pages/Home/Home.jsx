import {Button, Form, Input, message} from "antd";
import React from "react";
import './Home.css';

import 'antd/dist/antd.css'

function Home() {
    return <div>
        Hello, this is home page <a href={'/login'}>Login</a>
    </div>
}

export default Home;

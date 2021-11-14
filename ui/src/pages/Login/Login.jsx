import {Button, Form, Input, message} from "antd";
import React from "react";
import './Login.css';

import 'antd/dist/antd.css'
import {useState} from 'react';

function Login() {
    const [userName, setUserName] = useState('')
    const [password, setPassword] = useState('')
    const [loading, setLoading] = useState(false)

    const onUserChange  = (e)=> setUserName(e.target.value)
    const onPasswordChange = (e) => setPassword(e.target.value)

    function onLogin () {
        setLoading(true)
        fetch("/api/login").then((res) => {
            console.log(res);
        })
        setTimeout(() => {
            setLoading(false)
            message.success("Login is successful!")
        }, 2000)
    }

    return (
        <div className="App">
            <div className={'login-box'}>
                <Form
                    labelCol={{span: 5}}
                    wrapperCol={{span: 18}}
                >
                    <Form.Item label={'User Name'}>
                        <Input size={'large'} onChange={onUserChange} value={userName} />
                    </Form.Item>
                    <Form.Item label={'Password'}>
                        <Input size={'large'} onChange={onPasswordChange} value={password} type={'password'}/>
                    </Form.Item>
                    <Form.Item wrapperCol={{offset: 5, span: 16}}>
                        <Button  loading={loading} onClick={onLogin}  color={'red'} size={'long'} type={'primary'}>Login</Button>
                    </Form.Item>
                    <Form.Item>
                        <span>You are inputting:</span>
                        <br/>
                        <span style={{color: '#ccc'}}>user name:</span>
                        <span>{userName}</span>
                        <br/>
                        <span style={{color: '#ccc'}}>Password:</span>
                        <span>{password}</span>
                    </Form.Item>
                </Form>
            </div>
        </div>
    );
}

export default Login;

import {Button, Checkbox, Form, Input} from 'antd';
import axios from "axios";
import cookie from "react-cookies";
import {history} from "@@/core/history";

const Login = () => {
  const onFinish = (values: any) => {
    console.log('Success:', values);
    axios.post('/login', {username: values.username, password: values.password}).then(function (response) {
      console.log(response.data)
      cookie.save('token', response.data.data.token, {})
      history.push("/")
    }).catch(function (error) {
      console.log(error)
    })
  };

  const onFinishFailed = (errorInfo: any) => {
    console.log('Failed:', errorInfo);
  };

  return (
    <Form
      name="basic"
      labelCol={{span: 8}}
      wrapperCol={{span: 16}}
      initialValues={{remember: true}}
      onFinish={onFinish}
      onFinishFailed={onFinishFailed}
    >
      <Form.Item
        label="Username"
        name="username"
        rules={[{required: true, message: 'Please input your username!'}]}
      >
        <Input/>
      </Form.Item>

      <Form.Item
        label="Password"
        name="password"
        rules={[{required: true, message: 'Please input your password!'}]}
      >
        <Input.Password/>
      </Form.Item>

      <Form.Item name="remember" valuePropName="checked" wrapperCol={{offset: 8, span: 16}}>
        <Checkbox>Remember me</Checkbox>
      </Form.Item>

      <Form.Item wrapperCol={{offset: 8, span: 16}}>
        <Button type="primary" htmlType="submit">
          Submit
        </Button>
      </Form.Item>
    </Form>
  );
};

export default Login

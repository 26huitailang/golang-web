import React, { FC, ReactElement, useState } from 'react';
import AuthApi from '../../api/auth';

const Login: FC = (): ReactElement => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [message, setMessage] = useState('');
  const [display, setDisplay] = useState(false);

  const handleSubmit = (event: any) => {
    event.preventDefault();
    console.log(username, password);
    AuthApi.login({ username, password }).then((response) => {
      if (!response.success) {
        setDisplay(true);
        setMessage(response.details);
      }
    });
    // TODO 账户密码校验
    // todo 请求服务器，获取token，变更app的认证状态
  };

  const Hint: FC = (): ReactElement => (
    <span>
      Hint:
      {message}
    </span>
  );

  return (
    <div>
      {display ? <Hint /> : null}
      <form onSubmit={(event) => handleSubmit(event)}>
        <input
          type="text"
          className="username"
          value={username}
          onChange={(event) => setUsername(event.target.value)} />
        <input
          type="password"
          className="password"
          value={password}
          onChange={(event) => setPassword(event.target.value)} />
        <button type="submit">登录</button>
      </form>
    </div>
  );
};

export default Login;

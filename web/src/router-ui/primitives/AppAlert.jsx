import React from 'react';
import { Alert } from 'antd';

function AppAlert({ title, message, ...props }) {
  return <Alert {...props} title={title ?? message} />;
}

export default AppAlert;

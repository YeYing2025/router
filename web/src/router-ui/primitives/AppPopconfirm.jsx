import React from 'react';
import { Popconfirm } from 'antd';

function AppPopconfirm({ children, ...props }) {
  return <Popconfirm {...props}>{children}</Popconfirm>;
}

export default AppPopconfirm;

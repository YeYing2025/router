import React from 'react';
import { Layout } from 'antd';

function AppSider({ className = '', children, ...props }) {
  const nextClassName = ['router-ui-sider', className].filter(Boolean).join(' ');
  return (
    <Layout.Sider {...props} className={nextClassName}>
      {children}
    </Layout.Sider>
  );
}

export default AppSider;

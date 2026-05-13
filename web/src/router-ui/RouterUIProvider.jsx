import React from 'react';
import { ConfigProvider } from 'antd';
import zhCN from 'antd/locale/zh_CN';
import { antdTheme } from './theme/antd-theme';

function RouterUIProvider({ children }) {
  return (
    <ConfigProvider locale={zhCN} theme={antdTheme}>
      {children}
    </ConfigProvider>
  );
}

export default RouterUIProvider;

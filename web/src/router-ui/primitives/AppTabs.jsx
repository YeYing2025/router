import React from 'react';
import { Tabs } from 'antd';

function AppTabs({ className = '', items = [], ...props }) {
  const nextClassName = ['router-ui-tabs', className].filter(Boolean).join(' ');
  return <Tabs {...props} className={nextClassName} items={items} />;
}

export default AppTabs;

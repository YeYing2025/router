import React from 'react';
import { Card, Flex } from 'antd';

function AppSection({ className = '', children, title, extra }) {
  const nextClassName = ['router-ui-section', className].filter(Boolean).join(' ');
  return (
    <Card className={nextClassName} size='small'>
      {title || extra ? (
        <Flex
          className='router-ui-section-header'
          justify='space-between'
          align='center'
          wrap
        >
          <div className='router-ui-section-title'>{title}</div>
          {extra ? <div className='router-ui-section-extra'>{extra}</div> : null}
        </Flex>
      ) : null}
      <div className='router-ui-section-body'>{children}</div>
    </Card>
  );
}

export default AppSection;

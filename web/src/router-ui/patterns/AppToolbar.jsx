import React from 'react';
import { Flex } from 'antd';

function AppToolbar({
  className = '',
  start,
  end,
  startClassName = '',
  endClassName = '',
}) {
  const nextClassName = ['router-toolbar', className].filter(Boolean).join(' ');
  const nextStartClassName = ['router-toolbar-start', startClassName]
    .filter(Boolean)
    .join(' ');
  const nextEndClassName = ['router-toolbar-end', endClassName]
    .filter(Boolean)
    .join(' ');
  return (
    <Flex className={nextClassName} justify='space-between' align='center' wrap>
      <Flex className={nextStartClassName} align='center' wrap>
        {start}
      </Flex>
      <Flex className={nextEndClassName} align='center' wrap>
        {end}
      </Flex>
    </Flex>
  );
}

export default AppToolbar;

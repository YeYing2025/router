import React from 'react';
import { Card, Flex } from 'antd';

function AppDetailSection({
  className = '',
  title,
  titleTag: TitleTag = 'h3',
  titleClassName = 'router-entity-detail-section-title',
  headerStart,
  headerEnd,
  headerClassName = '',
  bodyClassName = '',
  children,
}) {
  const nextClassName = ['router-entity-detail-section', className]
    .filter(Boolean)
    .join(' ');
  const nextHeaderClassName = ['router-entity-detail-section-header', headerClassName]
    .filter(Boolean)
    .join(' ');
  const nextBodyClassName = [bodyClassName].filter(Boolean).join(' ');

  return (
    <Card className={nextClassName} size='small'>
      {(title || headerStart || headerEnd) ? (
        <Flex
          className={nextHeaderClassName}
          justify='space-between'
          align='center'
          wrap
        >
          <Flex className='router-toolbar-start' align='center' wrap>
            {title ? <TitleTag className={titleClassName}>{title}</TitleTag> : null}
            {headerStart}
          </Flex>
          {headerEnd ? (
            <Flex className='router-toolbar-end' align='center' wrap>
              {headerEnd}
            </Flex>
          ) : null}
        </Flex>
      ) : null}
      {nextBodyClassName ? (
        <div className={['router-entity-detail-section-body', nextBodyClassName].join(' ')}>
          {children}
        </div>
      ) : (
        <div className='router-entity-detail-section-body'>{children}</div>
      )}
    </Card>
  );
}

export default AppDetailSection;

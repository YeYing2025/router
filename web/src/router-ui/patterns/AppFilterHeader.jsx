import React from 'react';
import AppToolbar from './AppToolbar';

function AppFilterHeader({
  className = '',
  title,
  titleTag: TitleTag = 'div',
  titleClassName = 'router-toolbar-title',
  meta,
  metaClassName = 'router-toolbar-meta',
  actions,
  picker,
  query,
  startClassName = '',
  endClassName = '',
}) {
  const nextClassName = ['router-log-toolbar', 'router-block-gap-sm', className]
    .filter(Boolean)
    .join(' ');
  const nextStartClassName = ['router-log-toolbar-start', startClassName]
    .filter(Boolean)
    .join(' ');
  const resolvedStart =
    title || meta || picker ? (
      <>
        {title ? <TitleTag className={titleClassName}>{title}</TitleTag> : null}
        {meta ? <span className={metaClassName}>{meta}</span> : null}
        {picker}
      </>
    ) : null;
  const resolvedEnd =
    actions || query ? (
      <>
        {actions}
        {query}
      </>
    ) : null;

  return (
    <AppToolbar
      className={nextClassName}
      start={resolvedStart}
      end={resolvedEnd}
      startClassName={nextStartClassName}
      endClassName={endClassName}
    />
  );
}

export default AppFilterHeader;

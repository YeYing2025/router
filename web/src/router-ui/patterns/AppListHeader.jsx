import React from 'react';
import AppToolbar from './AppToolbar';

function AppListHeader({
  className = '',
  title,
  titleTag: TitleTag = 'div',
  titleClassName = 'router-toolbar-title',
  meta,
  metaClassName = 'router-toolbar-meta',
  actions,
  start,
  end,
  startClassName = '',
  endClassName = '',
}) {
  const nextClassName = ['router-block-gap-sm', className].filter(Boolean).join(' ');
  const resolvedStart =
    title || meta || start ? (
      <>
        {title ? <TitleTag className={titleClassName}>{title}</TitleTag> : null}
        {meta ? <span className={metaClassName}>{meta}</span> : null}
        {start}
      </>
    ) : null;
  const resolvedEnd =
    actions || end ? (
      <>
        {actions}
        {end}
      </>
    ) : null;
  return (
    <AppToolbar
      className={nextClassName}
      start={resolvedStart}
      end={resolvedEnd}
      startClassName={startClassName}
      endClassName={endClassName}
    />
  );
}

export default AppListHeader;

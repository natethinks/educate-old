import { EducatePage } from './app.po';

describe('educate App', () => {
  let page: EducatePage;

  beforeEach(() => {
    page = new EducatePage();
  });

  it('should display welcome message', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('Welcome to app!!');
  });
});

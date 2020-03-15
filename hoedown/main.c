#include "./src/buffer.h"
#include "./src/html.h"

void GenerateHTML(char *);

int main()
{
  printf("Start to hoedown test\n");
  GenerateHTML("- [ ] Test\n  - [x] tst");
  printf("Done\n");
  return 0;
}

void GenerateHTML(char *markdownText) {
  // variables
  FILE *fp;
  hoedown_buffer *ib, *ob;
  hoedown_renderer *renderer = NULL;
  hoedown_document *document;
  unsigned int extensions = 0;
  //unsigned int renderFlags = 0;
  // set extension && renderFlags
  extensions \
    = HOEDOWN_EXT_AUTOLINK \
    | HOEDOWN_EXT_FENCED_CODE \
    | HOEDOWN_EXT_FOOTNOTES \
    | HOEDOWN_EXT_HIGHLIGHT \
  //  | HOEDOWN_EXT_NO_INTRA_EMPHASIS
  //  | HOEDOWN_EXT_SPACE_HEADERS
    | HOEDOWN_EXT_QUOTE
    | HOEDOWN_EXT_STRIKETHROUGH \
    | HOEDOWN_EXT_SUPERSCRIPT \
    | HOEDOWN_EXT_TABLES \
    | HOEDOWN_EXT_UNDERLINE \
    | HOEDOWN_EXT_MATH \
    | HOEDOWN_EXT_MATH_EXPLICIT;
  // buffer size
  ib = hoedown_buffer_new(256);
  ob = hoedown_buffer_new(256);
  // set text
  hoedown_buffer_puts(ib, markdownText);
  // renderer
  renderer = hoedown_html_renderer_new(0, 0);
  //renderer = hoedown_html_toc_renderer_new(0);
  // document by render
  document = hoedown_document_new(renderer, 0, 16);
  // start render to markdown
  hoedown_document_render(document, ob, ib->data, ib->size);
  // save to test.html file
  fp = fopen("test.html", "w");
  (void)fwrite(ob->data, 1, ob->size, fp);
  // free all
  hoedown_buffer_free(ib);
  hoedown_buffer_free(ob);
  hoedown_html_renderer_free(renderer);
  hoedown_document_free(document);
  return;
}
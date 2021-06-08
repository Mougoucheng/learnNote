<center><font size=10 color="blue"> CORS  跨域资源共享 </font></center>



 

**[CORS](https://developer.mozilla.org/zh-CN/docs/Glossary/CORS)** （Cross-Origin Resource Sharing，`跨源资源共享` （或通俗地译为`跨域资源共享`））是一种基于[HTTP](https://developer.mozilla.org/zh-CN/docs/Glossary/HTTP) 头的机制，它由一系列传输的[HTTP头](https://developer.mozilla.org/zh-CN/docs/Glossary/HTTP_header)组成，该机制通过允许服务器标示除了它自己以外的其它[origin](https://developer.mozilla.org/zh-CN/docs/Glossary/Origin)（域，协议和端口），这样浏览器可以访问加载这些资源。跨源资源共享还通过一种机制来检查服务器是否会允许要发送的真实请求，该机制通过浏览器发起一个到服务器托管的跨源资源的"预检"请求。在预检中，浏览器发送的头中标示有HTTP方法和真实请求中会用到的头，这些HTTP头决定浏览器是否阻止前端 JavaScript 代码获取跨域请求的响应。

[同源安全策略](https://developer.mozilla.org/zh-CN/docs/Web/Security/Same-origin_policy) 默认阻止“跨域”获取资源。但是 CORS 给了web服务器这样的权限，即服务器可以选择，允许跨域请求访问到它们的资源。

# 同源政策

## 1.含义

1995年，同源政策由 Netscape 公司引入浏览器。目前，所有浏览器都实行这个政策。

最初，它的含义是指，A网页设置的 Cookie，B 网页不能打开，除非这两个网页"`同源`"。所谓 "`同源`"指的是三个相同:

* **协议相同**
* **域名相同**
* **端口相同**

> 举例来说，`http://www.example.com/dir/page.html` 这个网址，协议是 `http://`，域名是`www.example.com`，端口是`80`（默认端口可以省略）。它的同源情况如下：
>
> * http://www.example.com/dir2/other.html：同源
> * http://example.com/dir/other.html：不同源（域名不同）
> * http://v2.www.example.com/dir/other.html：不同源（域名不同）
> * http://www.example.com:81/dir2/other.html：不同源（端口不同）

## 2.目的

同源政策的目的，是为了保证用户信息的安全，防止恶意的网站窃取数据。

> 设想这样的一种情况：A网站是一家银行，用户登录以后，又去浏览其他网站。如果其他网站可以读取A网站的Cookie，会发生什么？很显然，如果Cookie包含隐私（比如存款总额），这些信息就会泄露。更可怕的是，Cookie往往用来保存用户的登录状态，如果用户没有退出登录，其他网站就可以冒充用户，为所欲为。因为浏览器同时还规定，提交表单不受同源政策的限制。

由此可见，`同源政策`是必需的，否则Cookie可以共享，互联网就毫无安全可言了。

## 3.限制范围

随着互联网的发展，`同源政策`越来越严格。目前，如果非同源，共有三种行为受到限制。

* Cookie、LocalStorage、IndexDB无法读取。
* DOM 无法获得。
* AJAX 请求不能发送。

虽然这些限制是必须的，但是有时很不方便，合理的用途也受到影响，如何规避上述三种限制，参考：[浏览器同源政策及其规避方法](http://www.ruanyifeng.com/blog/2016/04/same-origin-policy.html)，但是已不推荐，建议使用CORS（跨院资源共享，Cross-Origin Resource Sharing），它是W3C标准，是跨源AJAX请求的根本解决方法，CORS允许发送任何类型的请求。

# CORS

CORS是一个W3C标准，全称是 "跨域资源共享" （Cross-Origin Resource Sharing）。

它允许浏览器向跨源服务器，发出`XMLHttpRequest`请求，从而克服了AJAX只能同源使用的限制。

## 1.简介

CORS需要浏览器和服务器同时支持。目前，所有浏览器都支持该功能，IE浏览器不能低于IE10。

整个CORS通信过程，都是浏览器自动完成，不需要用户参与。对于开发者来说，CORS通信与同源的AJAX通信没有差别，代码完全一样。浏览器一旦发现AJAX请求跨源，就会自动添加一些附加的头信息，有时还会多出一次附加的请求，但用户不会有感觉。

因此，实现CORS通信的关键是服务器。只要服务器实现了CORS接口，就可以跨源通信。

## 2.两种请求

浏览器将CORS请求分成两类：

* 简单请求（simple request）
* 非简单请求（not-so-simple request）。

**`简单请求`需要满足以下条件：**

* 使用下列方法之一：
  * [`GET`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Methods/GET)
  * [`HEAD`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Methods/HEAD)
  * [`POST`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Methods/POST)
* HTTP头部信息只包含以下几种字段：
  * [`Accept`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Accept)
  * [`Accept-Language`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Accept-Language)
  * [`Content-Language`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Content-Language)
  * [`Content-Type`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Content-Type) （需要注意额外的限制）
  * `DPR`
  * `Downlink`
  * `Save-Data`
  * `Viewport-Width`
  * `Width`
* [`Content-Type`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Content-Type) 的值仅限于下列三者之一：
  * text / plain
  * multipart/form-data
  * application/x-www-form-urlencoded
* 请求中的任意`XMLHttpRequestUpload` 对象均没有注册任何事件监听器；`XMLHttpRequestUpload` 对象可以使用 [`XMLHttpRequest.upload`](https://developer.mozilla.org/zh-CN/docs/Web/API/XMLHttpRequest/upload) 属性访问。
* 请求中没有使用 [`ReadableStream`](https://developer.mozilla.org/zh-CN/docs/Web/API/ReadableStream) 对象。

这是为了兼容表单（form），因为历史上表单一直可以发出跨域请求。AJAX的跨域设计就是，只要表单可以发，AJAX就可以直接发。

其余凡是不同时满足上面两个条件的，就属于**`非简单请求`**。浏览器对于这两种请求的处理，是不一样的。

## 3.简单请求

### a).基本流程

对于简单请求，浏览器直接发出CORS请求。具体来说，就是在头信息之中，增加一个 `Origin` 字段。

![haha](https://mdn.mozillademos.org/files/17214/simple-req-updated.png)

下面是一个例子，浏览器发现这次跨源AJAX请求是简单请求，就自动在头信息之中，添加一个 `Origin` 字段。 

```http
GET /cors HTTP/1.1
Origin: http://api.bob.com
Host: api.alice.com
Accept-Language: en-US
Connection: keep-alive
User-Agent: Mozilla/5.0...
```

上面的头信息中，`Origin` 字段用来说明，本次请求来自哪个源（协议+域名+端口）。服务器根据这个值，决定是否统一这次请求。

如果 `Origin` 指定的源，不在服务器许可范围内，服务器会返回一个正常的HTTP响应。浏览器发现，这个响应的头信息没有包含 `Access-Control-Allow-Origin` 字段（详见下文），就知道出错了，从而抛出一个错误，被 `XMLHttpRequest` 的 `onerror` 回调函数捕获。注意，这种错误无法通过状态码识别，因为HTTP响应的状态码有可能是200。

如果 `Origin` 指定的域名在服务器许可范围内，服务器返回的响应中，会多出机构头信息字段。

```http
Access-Control-Allow-Origin: http://api.bob.com
Access-Control-Allow-Credentials: true
Access-Control-Expose-Headers: FooBar
Content-Type: text/html; charset=utf-8
```

上面的头信息之中，有三个与CORS请求相关的字段，都以 `Access-Control-` 开头。

* **Access-Control-Allow-Origin**

  该字段是必须的，它的值要么是请求时 `Origin` 字段的值，要么是一个 `*`，表示接受任意域名的请求。

* **Access-Control-Allow-Credentials**

  该字段可选。它的值是一个布尔值，表示是否允许发送Cookie。默认情况下，Cookie不包括在CORS请求之中。设为 `true` ，即表示服务器明确许可，Cookie可以包含在请求中，一起发给服务器。这个值也只能设为 `true` ，如果服务器不要浏览器发送Cookie，删除该字段即可。

* **Access-Control-Expose-Headers**

  该字段可选。CORS请求时，`XMLHttpRequest` 对象的 `getResponseHeader()` 方法只能拿到6个基本字段：`Cache-Control`、`Content-Language`、`Content-Type`、`Expires`、`Last-Modified`、`Pragma`。如果想拿到其他字段，就必须在 `Access-Control-Expose-Headers` 里面指定。上面的例子指定，`getResponseHeader('FooBar')` 可以返回 `FooBar` 字段的值。

### b).withCredentials属性

上面说到，CORS请求默认不发送Cookie和HTTP认证信息。如果要把Cookie发到服务器，一方面要服务器同意，指定 `Access-Control-Allow-Credentials` 字段。

```http
Access-Control-Allow-Credentials: true
```

另一方面，开发者必须在AJAX请求中打开 `withCredentials` 属性。

```javascript
var xhr = new XMLHttpRequest();
xhr.withCredentials = true;
```

否则，即使服务器同意发送Cookie，浏览器也不会发送。或者，服务器要求设置Cookie，浏览器也不会处理。

但是，如果省略 `withCredentials` 设置，有的浏览器还是会一起发送Cookie。这时，可以显示关闭  `withCredentials` 。

```javascript
xhr.withCredentials = false;
```

需要注意的是，如果要发送Cookie，`Access-Control-Allow-Origin` 就不能设为星号，必须指定明确的、与请求网页一致的域名。同时，Cookie依然遵循同源政策，只有用服务器域名设置的Cookie才会上传，其他域名的Cookie并不会上传，且（跨源）原网页代码中的 `document.cookie` 也无法读取服务器域名下的Cookie。

## 4.非简单请求

非简单请求是那种对服务器有特殊要求的请求，比如请求方法是 `PUT` 或 `DELETE`，或者 `Conten-Type` 字段的类型是 `application/json`，又或者HTTP请求头中有自定义字段，如请求头中新加appId、bizId字段用于发送数据。

### a).预检请求

非简单请求的CORS请求，会在正式通信之前，使用 `OPTIONS` 方法，增加一次HTTP查询请求发送到服务器，以获知服务器是否允许该实际请求，称为“预检”请求（preflight）。

浏览器先询问服务器，当前网页所在的域名是否在服务器的许可名单之中，以及可以使用哪些HTTP动词和头信息字段。只有得到肯定答复，浏览器才会发出正式的 `XMLHttpRequest` 请求，否则就报错。

下面是一段浏览器的JavaScrip脚本。

```javascript
var url = 'http://api.alice.com/cors';
var xhr = new XMLHttpRequest();
xhr.open('PUT', url, true);
xhr.setRequestHeader('X-Custom-Header', 'value');
xhr.send();
```

上面代码中，HTTP请求的方法是 `PUT` ，并且发送一个自定义的头信息 `X-Custom-Header`。

浏览器发现，这是一个非简单请求，就自动发出一个“预检”请求，要求服务器确认可以这样请求。下面是这个“预检”请求的HTTP头信息。

```http
OPTIONS /cors HTTP/1.1
Origin: http://api.bob.com
Access-Control-Request-Method: PUT
Access-Control-Request-Headers: X-Custom-Header
Host: api.alice.com
Accept-Language: en-US
Connection: keep-alive
User-Agent: Mozilla/5.0...
```

“预检”请求用的请求方法是 `OPTIONS`，表示这个请求是用来询问的。头信息里面，关键字段时 `Origin`，表示请求来自哪个源。

除了 `Origin` 字段，“预检”请求的头信息包括两个特殊字段。

* **Access-Control-Request-Method**

  该字段时必须的，用来列出浏览器的CORS请求会用到哪些HTTP方法，上例是 `PUT`。

* **Access-Control-Request-Headers**

  该字段是一个逗号分隔的字符串，指定浏览器CORS请求会额外发送的头信息字段，上例是 `X-Custom-Header`。

### b).预检请求的响应

服务器收到“预检”请求，检查了 `Origin`、`Access-Control-Request-Method` 和 `Access-Control-Request-Headers` 字段以后，确认允许跨源请求，就可以作出响应。

```http
HTTP/1.1 200 OK
Date: Mon, 01 Dec 2008 01:15:39 GMT
Server: Apache/2.0.61 (Unix)
Access-Control-Allow-Origin: http://api.bob.com
Access-Control-Allow-Methods: GET, POST, PUT
Access-Control-Allow-Headers: X-Custom-Header
Content-Type: text/html; charset=utf-8
Content-Encoding: gzip
Content-Length: 0
Keep-Alive: timeout=2, max=100
Connection: Keep-Alive
Content-Type: text/plain
```

上面HTTP响应中，关键的是 `Access-Control-Allow-Origin` 字段，表示 `http://api.bob.com` 可以请求数据。该字段也可以设为星号，表示同意任意跨源请求。

```http
Access-Control-Allow-Origin: *
```

如果服务器否定了“预检”请求，会返回一个正常的HTTP响应，但是没有任何CORS相关的头信息字段。这时，浏览器就会认定，服务器不同意预检请求，因此触发一个错误，被 `XMLHttpRequest` 对象的 `onerror` 回调函数捕获。控制台会打印出如下的报错信息。

```http
XMLHttpRequest cannot load http://api.alice.com.
Origin http://api.bob.com is not allowed by Access-Control-Allow-Origin.
```

服务器响应的其他CORS相关字段如下。

```http
Access-Control-Allow-Methods: GET, POST, PUT
Access-Control-Allow-Headers: X-Custom-Header
Access-Control-Allow-Credentials: true
Access-Control-Max-Age: 1728000
```

* **Access-Control-Allow-Methods**

  该字段必需，它的值是英文逗号分隔的一个字符串，表明服务器支持的所有跨域请求的方法。注意，返回的是所有支持的方法，而不单是浏览器请求的那个方法。这是为了避免多次“预检”请求。

* **Access-Control-Allow-Headers**

  如果浏览器请求包括 `Access-Control-Request-Headers` 字段，则 `Access-Control-Allow-Headers` 字段是必需的。它也是一个英文逗号分隔的字符串，表明服务器支持的所有头信息字段，不限于浏览器在“预检”中请求的字段。

* **Access-Control-Allow-Credentials**

  该字段与简单请求时的含义相同。

* **Access-Control-Max-Age**

  该字段可选，用来指定本次预检请求的有效期，单位为秒。上面结果中，有效期是20天（1728000秒），即允许缓存该条响应20天，在此期间，不用发出另一条预检请求。

  

### c).浏览器的正常请求和响应

  一旦服务器通过了“预检”请求，以后每次浏览器正常的CORS请求，就都跟简单请求一样，会有一个 `Origin` 头信息字段。服务器的响应，也都会有一个 `Access-Control-Allow-Origin` 头信息。

下面是“预检”请求之后，浏览器的正常CORS请求。

```http
PUT /cors HTTP/1.1
Origin: http://api.bob.com
Host: api.alice.com
X-Custom-Header: value
Accept-Language: en-US
Connection: keep-alive
User-Agent: Mozilla/5.0...
```

上面头信息的 `Origin` 字段是浏览器自动添加的。

下面是服务器正常的响应。

```http
Access-Control-Allow-Origin: http://api.bob.com
Content-Type: text/html; charset=utf-8
```

上面头信息中，`Access-Control-Allow-Origin` 字段是每次响应都必定包含的。

# CORS与JSONP的比较

CORS与JSONP的使用目的相同，都是为了解决浏览器跨源问题。

JSONP优势:

* 只支持 `GET` 请求
* 支持老式浏览器
* 可以向不支持CORS的网站请求数据

CORS优势:

* 支持所有类型的HTTP请求
* 比JSONP更强大。

# 总结

CORS，全称：跨源资源共享（Cross-Origin Resource Sharing），由于同源政策导致浏览器请求非同源资源时报错，通过“预检”请求（非简单请求）查询服务器，确认获得服务器许可，从而解决跨源资源共享问题。（只是使用浏览器时会出现，代码请求时不会出现CORS问题）

头信息字段列表：

> **`请求：`**
>
> * **Origin**
> * **Access-Control-Request-Method**
> * **Access-Control-Request-Headers**
>
> (请注意，这些首部字段无须手动设置。 当开发者使用 XMLHttpRequest 对象发起跨源请求时，它们已经被设置就绪。)

> **`响应：`**
>
> * **Access-Control-Allow-Origin**
> * **Access-Control-Allow-Methods**
> * **Access-Control-Allow-Headers**
> * **Access-Control-Allow-Credentials**
> * **Access-Control-Expose-Headers**
> * **Access-Control-Max-Age**

# 参考

* [CORS](https://developer.mozilla.org/zh-CN/docs/Glossary/CORS)

* [跨源资源共享（CORS）](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/CORS)

* [阮一峰：跨域资源共享 CORS 详解](http://www.ruanyifeng.com/blog/2016/04/cors.html)

  


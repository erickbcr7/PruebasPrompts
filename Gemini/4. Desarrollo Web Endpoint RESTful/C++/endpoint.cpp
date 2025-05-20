#include <boost/network/http/server.hpp>
#include <iostream>
#include <string>
#include <map>

namespace http = boost::network::http;

struct saludo_responder;
typedef http::server<saludo_responder> server;

struct saludo_responder {
    void operator()(server::request const& request, server::response& response) {
        std::map<std::string, std::string> query_params;
        std::string query = request.query();
        size_t pos = 0;
        while (pos < query.length()) {
            size_t amp_pos = query.find('&', pos);
            std::string pair = query.substr(pos, amp_pos == std::string::npos ? std::string::npos : amp_pos - pos);
            size_t eq_pos = pair.find('=');
            if (eq_pos != std::string::npos) {
                query_params[pair.substr(0, eq_pos)] = pair.substr(eq_pos + 1);
            }
            if (amp_pos == std::string::npos) break;
            pos = amp_pos + 1;
        }

        auto it = query_params.find("nombre");
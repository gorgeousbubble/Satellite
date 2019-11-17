using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;
using System.IO;
using System.Net;
using System.Threading;
using Newtonsoft.Json;
using Newtonsoft.Json.Linq;

namespace qrcode
{
    public partial class FormMain : Form
    {
        public FormMain()
        {
            InitializeComponent();
        }

        private void FormMain_Load(object sender, EventArgs e)
        {
            // FormMain Initial
            AutoScaleMode = AutoScaleMode.Dpi;
            MaximizeBox = false;
            MinimizeBox = true;
            DoubleBuffered = true;
            StartPosition = FormStartPosition.CenterScreen;

            // QRCode
            this.textBox_images.ReadOnly = true;
        }

        private void FormMain_FormClosed(object sender, FormClosedEventArgs e)
        {

        }

        private string Get_json()
        {
            // check content file
            if (textBox_content.Text == "")
            {
                MessageBox.Show("Please input content!", "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                return "";
            }

            string content = string.Format("\"{0}\"", textBox_content.Text);

            string body = "";
            body += "{";
            body += string.Format("\"content\":{0},\"size\":256", content);
            body += "}";
            body = body.Replace('\\', '/');
            return body;
        }

        private void Save_image(string path, string content)
        {
            using (StreamWriter writer = new StreamWriter(path))
            {
                byte[] imageBytes = Encoding.ASCII.GetBytes(content);
                writer.Write(imageBytes);
            }
        }

        private void Button_generate_Click(object sender, EventArgs e)
        {
            try
            {
                string body = Get_json();
                if (body == "")
                    return;

                string path = textBox_images.Text;
                if (path == "")
                {
                    MessageBox.Show("Please input save path!", "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                    return;
                }

                byte[] requestBody = Encoding.ASCII.GetBytes(body);

                HttpWebRequest request = (HttpWebRequest)WebRequest.Create("http://localhost:8080/satellite/images/qrcode");
                request.Method = "POST";
                request.KeepAlive = false;
                request.ProtocolVersion = HttpVersion.Version11;
                request.ContentType = "application/json";
                request.ContentLength = requestBody.Length;
                using (Stream reqStream = request.GetRequestStream())
                {
                    reqStream.Write(requestBody, 0, requestBody.Length);
                }
                using (HttpWebResponse response = (HttpWebResponse)request.GetResponse())
                {
                    if (response.StatusCode == HttpStatusCode.OK)
                    {
                        string responseContent = "";
                        using (Stream resStream = response.GetResponseStream())
                        {
                            using (StreamReader reader = new StreamReader(resStream, Encoding.Default))
                            {
                                responseContent = reader.ReadToEnd().ToString();
                                // parser json struct...
                                //JObject jObject = (JObject)JsonConvert.DeserializeObject(responseContent);
                                byte[] imageBytes = Encoding.Default.GetBytes(responseContent);
                                //Save_image(path, responseContent);
                                // show qrcode images...
                                MemoryStream memoryStream = new MemoryStream(imageBytes, 0, imageBytes.Length);
                                Image image = Image.FromStream(memoryStream);
                                this.pictureBox_qrcode.SizeMode = PictureBoxSizeMode.StretchImage;
                                this.pictureBox_qrcode.Image = image;
                            }
                        }
                    }
                }
            }
            catch(Exception ex)
            {
                MessageBox.Show(ex.Message, "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
            }   
        }

        private void Button_clear_Click(object sender, EventArgs e)
        {

        }

        private void Button_select_Click(object sender, EventArgs e)
        {
            SaveFileDialog dialog = new SaveFileDialog();
            dialog.Filter = "All files(*.*)|*.*";
            dialog.FileName = "undefine.png";
            dialog.RestoreDirectory = true;

            DialogResult dr = dialog.ShowDialog();
            if (dr == DialogResult.OK)
            {
                textBox_images.Text = dialog.FileName;
            }
        }
    }
}

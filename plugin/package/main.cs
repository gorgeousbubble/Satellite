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

namespace package
{
    public partial class FormMain : Form
    {
        public class T_PackInfo
        {
            public int Number { get; set; }
            public string Name { get; set; }
            public string Path { get; set; }
        }

        public class T_UnpackInfo
        {
            public int Number { get; set; }
            public string Name { get; set; }
            public string Size { get; set; }
            public string Type { get; set; }
        }

        private Thread m_tPack;
        private Thread m_tUnpack;
        private Thread m_tVerbose;
        private List<T_PackInfo> m_vecPackInfo = new List<T_PackInfo>();
        private List<T_UnpackInfo> m_vecUnpackInfo = new List<T_UnpackInfo>();

        private delegate void UPDATELISTVIEWCALLBACK();
        private UPDATELISTVIEWCALLBACK UpdateListViewCallback;

        private bool m_bChoose = false;

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

            // TabPack Initial
            listView_pack.View = View.Details;
            listView_pack.CheckBoxes = true;
            listView_pack.Columns.Add("number", 60, HorizontalAlignment.Center);
            listView_pack.Columns.Add("files", 240, HorizontalAlignment.Center);
            listView_pack.Columns.Add("detials", 400, HorizontalAlignment.Center);

            comboBox_pack.DropDownStyle = ComboBoxStyle.DropDownList;
            comboBox_pack.Items.Clear();
            comboBox_pack.Items.Add("aes");
            comboBox_pack.Items.Add("3des");
            comboBox_pack.Items.Add("des");
            comboBox_pack.Items.Add("rsa");
            comboBox_pack.Items.Add("base64");
            comboBox_pack.SelectedIndex = 0;

            textBox_pack_path.ReadOnly = true;

            timer_pack_process.Enabled = true;
            timer_pack_process.Stop();

            // TabUnpack Initial
            listView_unpack.View = View.Details;
            listView_unpack.Columns.Add("number", 60, HorizontalAlignment.Center);
            listView_unpack.Columns.Add("files", 240, HorizontalAlignment.Center);
            listView_unpack.Columns.Add("size", 200, HorizontalAlignment.Center);
            listView_unpack.Columns.Add("type", 200, HorizontalAlignment.Center);

            textBox_unpack_src.ReadOnly = true;
            textBox_unpack_dest.ReadOnly = true;

            UpdateListViewCallback = new UPDATELISTVIEWCALLBACK(Update_unpack_listView);
        }

        private void FormMain_FormClosed(object sender, FormClosedEventArgs e)
        {

        }

        private void Update_pack_listView()
        {
            listView_pack.Items.Clear();
            listView_pack.BeginUpdate();
            for (int i = 0; i < m_vecPackInfo.Count; ++i)
            {
                ListViewItem item = new ListViewItem();
                item.Text = m_vecPackInfo[i].Number.ToString();
                item.SubItems.Add(m_vecPackInfo[i].Name);
                item.SubItems.Add(m_vecPackInfo[i].Path);
                listView_pack.Items.Add(item);
            }
            listView_pack.EndUpdate();
        }

        private void Update_unpack_listView()
        {
            listView_unpack.Items.Clear();
            listView_unpack.BeginUpdate();
            for (int i = 0; i < m_vecUnpackInfo.Count; ++i)
            {
                ListViewItem item = new ListViewItem();
                item.Text = m_vecUnpackInfo[i].Number.ToString();
                item.SubItems.Add(m_vecUnpackInfo[i].Name);
                item.SubItems.Add(m_vecUnpackInfo[i].Size);
                item.SubItems.Add(m_vecUnpackInfo[i].Type);
                listView_unpack.Items.Add(item);
            }
            listView_unpack.EndUpdate();
        }

        private string Get_pack_json()
        {
            // check src file list
            if (m_vecPackInfo.Count <= 0)
            {
                MessageBox.Show("Please select src files name!", "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                return "";
            }
            // check dest file
            if (textBox_pack_path.Text == "")
            {
                MessageBox.Show("Please select dest file path!", "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                return "";
            }

            string src = "";
            for (int i = 0; i < m_vecPackInfo.Count - 1; ++i)
            {
                src += string.Format("\"{0}\",", m_vecPackInfo[i].Path);
            }
            src += string.Format("\"{0}\"", m_vecPackInfo[m_vecPackInfo.Count - 1].Path);

            string dest = string.Format("\"{0}\"", textBox_pack_path.Text);

            string type = "";
            switch (comboBox_pack.SelectedIndex)
            {
                case 0:
                    type = "\"aes\"";
                    break;
                case 1:
                    type = "\"3des\"";
                    break;
                case 2:
                    type = "\"des\"";
                    break;
                case 3:
                    type = "\"rsa\"";
                    break;
                case 4:
                    type = "\"base64\"";
                    break;
                default:
                    type = "\"aes\"";
                    break;
            }

            string body = "";
            body += "{";
            body += string.Format("\"src\":[{0}],\"dest\":{1},\"type\":{2}", src, dest, type);
            body += "}";
            body = body.Replace('\\', '/');
            return body;
        }

        private string Get_pack_process_json()
        {
            // check src file list
            if (m_vecPackInfo.Count <= 0)
            {
                MessageBox.Show("Please select src files name!", "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                return "";
            }
            // check dest file
            if (textBox_pack_path.Text == "")
            {
                MessageBox.Show("Please select dest file path!", "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                return "";
            }

            string src = "";
            for (int i = 0; i < m_vecPackInfo.Count - 1; ++i)
            {
                src += string.Format("\"{0}\",", m_vecPackInfo[i].Path);
            }
            src += string.Format("\"{0}\"", m_vecPackInfo[m_vecPackInfo.Count - 1].Path);

            string type = "";
            switch (comboBox_pack.SelectedIndex)
            {
                case 0:
                    type = "\"aes\"";
                    break;
                case 1:
                    type = "\"3des\"";
                    break;
                case 2:
                    type = "\"des\"";
                    break;
                case 3:
                    type = "\"rsa\"";
                    break;
                case 4:
                    type = "\"base64\"";
                    break;
                default:
                    type = "\"aes\"";
                    break;
            }

            string body = "";
            body += "{";
            body += string.Format("\"src\":[{0}],\"type\":{1}", src, type);
            body += "}";
            body = body.Replace('\\', '/');
            return body;
        }

        private string Get_unpack_json()
        {
            // check src file
            if (textBox_unpack_src.Text == "")
            {
                MessageBox.Show("Please select src file name!", "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                return "";
            }
            // check dest path
            if (textBox_unpack_dest.Text == "")
            {
                MessageBox.Show("Please select dest file path!", "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                return "";
            }

            string src = string.Format("\"{0}\"", textBox_unpack_src.Text);
            string dest = string.Format("\"{0}\"", textBox_unpack_dest.Text);

            string body = "";
            body += "{";
            if (m_bChoose == true && listView_unpack.CheckedItems.Count > 0)
            {
                string target = string.Format("\"{0}\"", m_vecUnpackInfo[listView_unpack.CheckedItems[0].Index].Name);
                body += string.Format("\"src\":{0},\"target\":{1},\"dest\":{2}", src, target, dest);
            }
            else
            {
                body += string.Format("\"src\":{0},\"dest\":{1}", src, dest);
            }
            body += "}";
            body = body.Replace('\\', '/');
            return body;
        }

        private string Get_unpack_verbose_json()
        {
            // check src file
            if (textBox_unpack_src.Text == "")
            {
                MessageBox.Show("Please select src file name!", "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                return "";
            }

            string src = string.Format("\"{0}\"", textBox_unpack_src.Text);

            string body = "";
            body += "{";
            body += string.Format("\"src\":{0}", src);
            body += "}";
            body = body.Replace('\\', '/');
            return body;
        }

        private string Get_unpack_process_json()
        {
            // check src file
            if (textBox_unpack_src.Text == "")
            {
                MessageBox.Show("Please select src file name!", "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                return "";
            }

            string src = string.Format("\"{0}\"", textBox_unpack_src.Text);

            string body = "";
            body += "{";
            body += string.Format("\"src\":{0}", src);
            body += "}";
            body = body.Replace('\\', '/');
            return body;
        }

        private void PackThread(object sender)
        {
            string body = sender as string;
            try
            {
                byte[] requestBody = Encoding.ASCII.GetBytes(body);

                HttpWebRequest request = (HttpWebRequest)WebRequest.Create("http://localhost:8080/satellite/pack");
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
                    if (response.StatusCode != HttpStatusCode.OK)
                    {
                        switch (response.StatusCode)
                        {
                            case HttpStatusCode.BadRequest:
                                MessageBox.Show("Bad Request(400)!", "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                                break;
                            case HttpStatusCode.InternalServerError:
                                MessageBox.Show("Internal Server Error(500)!", "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                                break;
                            default:
                                MessageBox.Show("Unknown Error!", "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                                break;
                        }
                        return;
                    }
                    MessageBox.Show("Pack Success!", "Information", MessageBoxButtons.OK, MessageBoxIcon.Information);
                }
            }
            catch (Exception ex)
            {
                MessageBox.Show(ex.Message, "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
            }
        }

        private void UnpackThread(object sender)
        {
            string body = sender as string;
            try
            {
                byte[] requestBody = Encoding.ASCII.GetBytes(body);
                string url = "";
                if (m_bChoose == true)
                {
                    url = "http://localhost:8080/satellite/unpack/f";
                }
                else
                {
                    url = "http://localhost:8080/satellite/unpack";
                }

                HttpWebRequest request = (HttpWebRequest)WebRequest.Create(url);
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
                    if (response.StatusCode != HttpStatusCode.OK)
                    {
                        switch (response.StatusCode)
                        {
                            case HttpStatusCode.BadRequest:
                                MessageBox.Show("Bad Request(400)!", "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                                break;
                            case HttpStatusCode.InternalServerError:
                                MessageBox.Show("Internal Server Error(500)!", "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                                break;
                            default:
                                MessageBox.Show("Unknown Error!", "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                                break;
                        }
                        return;
                    }
                    MessageBox.Show("Unpack Success!", "Information", MessageBoxButtons.OK, MessageBoxIcon.Information);
                }
            }
            catch (Exception ex)
            {
                MessageBox.Show(ex.Message, "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
            }
        }

        private void UnpackVerboseThread(object sender)
        {
            string body = sender as string;
            try
            {
                byte[] requestBody = Encoding.ASCII.GetBytes(body);

                HttpWebRequest request = (HttpWebRequest)WebRequest.Create("http://localhost:8080/satellite/unpack/v");
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
                    if (response.StatusCode != HttpStatusCode.OK)
                    {
                        switch (response.StatusCode)
                        {
                            case HttpStatusCode.BadRequest:
                                MessageBox.Show("Bad Request(400)!", "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                                break;
                            case HttpStatusCode.InternalServerError:
                                MessageBox.Show("Internal Server Error(500)!", "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                                break;
                            default:
                                MessageBox.Show("Unknown Error!", "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                                break;
                        }
                        return;
                    }

                    string responseContent = "";
                    using (Stream resStream = response.GetResponseStream())
                    {
                        using (StreamReader reader = new StreamReader(resStream, Encoding.UTF8))
                        {
                            responseContent = reader.ReadToEnd().ToString();
                            // parser json struct...
                            m_vecUnpackInfo.Clear();
                            JObject jFiles = (JObject)JsonConvert.DeserializeObject(responseContent);
                            JArray jArray = JArray.Parse(jFiles["files"].ToString());
                            foreach (var jItem in jArray)
                            {
                                JObject jObject = (JObject)jItem;
                                T_UnpackInfo info = new T_UnpackInfo();
                                info.Number = m_vecUnpackInfo.Count + 1;
                                info.Name = jObject["name"].ToString();
                                info.Size = jObject["size"].ToString();
                                info.Type = jObject["type"].ToString();
                                m_vecUnpackInfo.Add(info);
                            }
                        }
                    }
                    listView_unpack.Invoke(UpdateListViewCallback);
                }
            }
            catch (Exception ex)
            {
                MessageBox.Show(ex.Message, "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
            }
        }

        private void Button_pack_add_Click(object sender, EventArgs e)
        {
            OpenFileDialog dialog = new OpenFileDialog();
            dialog.Filter = "All files(*.*)|*.*";
            dialog.Multiselect = true;
            dialog.RestoreDirectory = true;

            DialogResult dr = dialog.ShowDialog();
            if (dr == DialogResult.OK)
            {
                foreach (var i in dialog.FileNames)
                {
                    bool flag = false;
                    foreach (var j in m_vecPackInfo)
                    {
                        if (j.Path == i)
                        {
                            flag = true;
                            break;
                        }
                    }
                    if (!flag)
                    {
                        T_PackInfo packInfo = new T_PackInfo();
                        FileInfo file = new FileInfo(i);
                        packInfo.Number = m_vecPackInfo.Count + 1;
                        packInfo.Name = file.Name;
                        packInfo.Path = i;
                        m_vecPackInfo.Add(packInfo);
                    }
                }

                Update_pack_listView();
            }

        }

        private void Button_pack_delete_Click(object sender, EventArgs e)
        {
            if (listView_pack.CheckedItems.Count > 0)
            {
                for (int i = 0; i < listView_pack.CheckedItems.Count; ++i)
                {
                    for (int j = 0; j < m_vecPackInfo.Count;)
                    {
                        if (m_vecPackInfo[j].Number.ToString() == listView_pack.CheckedItems[i].Text)
                        {
                            m_vecPackInfo.RemoveAt(j);
                        }
                        else
                        {
                            ++j;
                        }
                    }
                }

                for (int i = 0; i < m_vecPackInfo.Count; ++i)
                {
                    m_vecPackInfo[i].Number = i + 1;
                }

                Update_pack_listView();
            }
        }

        private void Button_pack_select_Click(object sender, EventArgs e)
        {
            SaveFileDialog dialog = new SaveFileDialog();
            dialog.Filter = "All files(*.*)|*.*";
            dialog.FileName = "undefine.pak";
            dialog.RestoreDirectory = true;

            DialogResult dr = dialog.ShowDialog();
            if (dr == DialogResult.OK)
            {
                textBox_pack_path.Text = dialog.FileName;
            }

        }

        private void Button_pack_execute_Click(object sender, EventArgs e)
        {
            string body = Get_pack_json();
            if (body == "")
                return;

            m_tPack = new Thread(new ParameterizedThreadStart(PackThread));
            m_tPack.IsBackground = true;
            m_tPack.Start(body);
            timer_pack_process.Start();
        }

        private void Timer_pack_process_Tick(object sender, EventArgs e)
        {
            try
            {
                string body = Get_pack_process_json();
                if (body == "")
                    return;

                byte[] requestBody = Encoding.ASCII.GetBytes(body);

                HttpWebRequest request = (HttpWebRequest)WebRequest.Create("http://localhost:8080/satellite/pack/p");
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
                            using (StreamReader reader = new StreamReader(resStream, Encoding.UTF8))
                            {
                                responseContent = reader.ReadToEnd().ToString();
                                // parser json struct...
                                JObject jObject = (JObject)JsonConvert.DeserializeObject(responseContent);
                                Int64 done = Convert.ToInt64(jObject["done"].ToString());
                                Int64 work = Convert.ToInt64(jObject["work"].ToString()) / 128;

                                int value = (int)((float)done * 100 / (float)work);
                                if (value >= 100)
                                {
                                    value = 100;
                                    timer_pack_process.Stop();
                                }
                                progressBar_pack.Value = value;
                            }
                        }
                    }
                }
            }
            catch (Exception ex)
            {
                timer_pack_process.Stop();
                MessageBox.Show(ex.Message, "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
            }
        }

        private void Button_unpack_src_Click(object sender, EventArgs e)
        {
            OpenFileDialog dialog = new OpenFileDialog();
            dialog.Filter = "All files(*.*)|*.*";
            dialog.Multiselect = false;
            dialog.RestoreDirectory = true;

            DialogResult dr = dialog.ShowDialog();
            if (dr == DialogResult.OK)
            {
                textBox_unpack_src.Text = dialog.FileName;
            }
        }

        private void Button_unpack_verbose_Click(object sender, EventArgs e)
        {
            string body = Get_unpack_verbose_json();
            if (body == "")
                return;

            m_tVerbose = new Thread(new ParameterizedThreadStart(UnpackVerboseThread));
            m_tVerbose.IsBackground = true;
            m_tVerbose.Start(body);
        }

        private void Button_unpack_dest_Click(object sender, EventArgs e)
        {
            SaveFileDialog dialog = new SaveFileDialog();
            dialog.Filter = "All files(*.*)|*.*";
            dialog.FileName = "undefine";
            dialog.RestoreDirectory = true;

            DialogResult dr = dialog.ShowDialog();
            if (dr == DialogResult.OK)
            {
                textBox_unpack_dest.Text = Path.GetDirectoryName(dialog.FileName) + "\\";
            }
        }

        private void Button_unpack_execute_Click(object sender, EventArgs e)
        {
            string body = Get_unpack_json();
            if (body == "")
                return;

            m_tUnpack = new Thread(new ParameterizedThreadStart(UnpackThread));
            m_tUnpack.IsBackground = true;
            m_tUnpack.Start(body);
            timer_unpack_process.Start();
        }

        private void Timer_unpack_process_Tick(object sender, EventArgs e)
        {
            try
            {
                string body = Get_unpack_process_json();
                if (body == "")
                    return;

                byte[] requestBody = Encoding.ASCII.GetBytes(body);

                HttpWebRequest request = (HttpWebRequest)WebRequest.Create("http://localhost:8080/satellite/unpack/p");
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
                            using (StreamReader reader = new StreamReader(resStream, Encoding.UTF8))
                            {
                                responseContent = reader.ReadToEnd().ToString();
                                // parser json struct...
                                JObject jObject = (JObject)JsonConvert.DeserializeObject(responseContent);
                                Int64 done = Convert.ToInt64(jObject["done"].ToString());
                                Int64 work = Convert.ToInt64(jObject["work"].ToString()) / 128;

                                int value = (int)((float)done * 100 / (float)work);
                                if (value >= 100)
                                {
                                    value = 100;
                                    timer_unpack_process.Stop();
                                }
                                progressBar_unpack.Value = value;
                            }
                        }
                    }
                }
            }
            catch (Exception ex)
            {
                timer_unpack_process.Stop();
                MessageBox.Show(ex.Message, "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
            }
        }

        private void Button_unpack_choose_Click(object sender, EventArgs e)
        {
            if (m_bChoose == false)
            {
                m_bChoose = true;
                listView_unpack.CheckBoxes = true;
                button_unpack_choose.Text = "Cancel";
            }
            else
            {
                m_bChoose = false;
                listView_unpack.CheckBoxes = false;
                button_unpack_choose.Text = "Choose";
            }
        }

        private void ListView_unpack_ItemChecked(object sender, ItemCheckedEventArgs e)
        {
            if (e.Item.Checked)
            {
                foreach (ListViewItem item in listView_unpack.CheckedItems)
                {
                    if (item != e.Item)
                    {
                        item.Checked = false;
                    }
                }
            }
        }
    }
}
